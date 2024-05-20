package interactor

import (
	"context"

	"github.com/nao1215/honeycomb/app/model"
	"github.com/nao1215/honeycomb/app/service"
	"github.com/nao1215/honeycomb/app/usecase"
)

var _ usecase.FollowLister = (*FollowLister)(nil)

// FollowLister is the interactor that wraps the basic FollowLister method.
type FollowLister struct {
	service.EventsLister
}

// NewFollowLister is the constructor of the FollowLister interactor.
func NewFollowLister(
	eventsLister service.EventsLister,
) *FollowLister {
	return &FollowLister{
		EventsLister: eventsLister,
	}
}

// ListFollow lists the followers of the user.
func (f *FollowLister) ListFollow(ctx context.Context, input *usecase.FollowListerInput) (*usecase.FollowListerOutput, error) {
	eventsOutput, err := f.EventsLister.ListEvents(ctx, &service.EventsListerInput{
		Filter: model.MyFollowFilter(input.PublicKey),
		Relay:  input.ConnectedRelay,
	})
	if err != nil {
		return nil, err
	}

	follows := make([]model.PublicKey, 0)
	for _, event := range eventsOutput.Events {
		for _, tag := range event.Tags {
			// The p tag, used to refer to another user: ["p", <32-bytes lowercase hex of a pubkey>, <recommended relay URL, optional>]
			// See https://github.com/nostr-protocol/nips/blob/master/01.md
			if len(tag) >= 2 && tag[0] == "p" {
				follows = append(follows, model.PublicKey(tag[1]))
			}
		}
	}

	if len(follows) == 0 {
		return &usecase.FollowListerOutput{
			Follows: []*model.Follow{},
		}, nil
	}

	followsProfile := make([]*model.Follow, 0, len(follows))
	for i := 0; i < len(follows); i += 100 {
		end := i + 100
		if end > len(follows) {
			end = len(follows)
		}

		eo, err := f.EventsLister.ListEvents(ctx, &service.EventsListerInput{
			Filter: model.ProfilesFilter(follows[i:end]),
			Relay:  input.ConnectedRelay,
		})
		if err != nil {
			return nil, err
		}

		profiles, err := toProfiles(eo.Events)
		if err != nil {
			return nil, err
		}
		for i, profile := range profiles {
			followsProfile = append(followsProfile, &model.Follow{
				PublicKey: follows[i], // follows[i] is the public key of the user.
				Profile:   *profile,
			})
		}
	}
	return &usecase.FollowListerOutput{
		Follows: followsProfile,
	}, nil
}
