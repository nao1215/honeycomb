// Package interactor contains the implementations of usecases.
package interactor

import (
	"context"
	"encoding/json"

	"github.com/nao1215/honeycomb/app/model"
	"github.com/nao1215/honeycomb/app/service"
	"github.com/nao1215/honeycomb/app/usecase"
	"github.com/nbd-wtf/go-nostr"
)

var _ usecase.ProfileGetter = (*ProfileGetter)(nil)

// ProfileGetter implements the ProfileGetter interface.
type ProfileGetter struct {
	service.RelayFinder
	service.EventsLister
}

// NewProfileGetter creates a new ProfileGetter.
func NewProfileGetter(
	relayFinder service.RelayFinder,
	profileGetter service.EventsLister,
) *ProfileGetter {
	return &ProfileGetter{
		RelayFinder:  relayFinder,
		EventsLister: profileGetter,
	}
}

// GetProfile gets the user's profile.
func (p *ProfileGetter) GetProfile(ctx context.Context, input *usecase.ProfileGetterInput) (*usecase.ProfileGetterOutput, error) {
	// TODO: Get Relays data fron upper layer.
	relayOutput, err := p.RelayFinder.FindRelay(ctx, &service.RelayFinderInput{
		Relays: map[model.WSS]model.Relay{
			"wss://relay-jp.nostr.wirednet.jp": {
				WSS:    "wss://relay-jp.nostr.wirednet.jp",
				Read:   true,
				Write:  true,
				Search: true,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	defer relayOutput.Relay.Close() //nolint:errcheck

	publicKey, err := input.NsecretKey.ToPublicKey()
	if err != nil {
		return nil, err
	}

	eventsOutput, err := p.EventsLister.ListEvents(ctx, &service.EventsListerInput{
		Filter: p.profileFilter(publicKey),
		Relay:  relayOutput.Relay,
	})
	if err != nil {
		return nil, err
	}
	return p.toProfileGetterOutput(eventsOutput.Events, publicKey)
}

// profileFilter returns the filter for getting the user's profile.
func (p *ProfileGetter) profileFilter(publicKey model.PublicKey) nostr.Filter {
	return nostr.Filter{
		Kinds:   []int{nostr.KindProfileMetadata},
		Authors: []string{publicKey.String()},
		Limit:   1,
	}
}

// toProfileGetterOutput converts the events to the ProfileGetterOutput.
func (p *ProfileGetter) toProfileGetterOutput(events []*nostr.Event, publicKey model.PublicKey) (*usecase.ProfileGetterOutput, error) {
	if len(events) == 0 {
		return nil, usecase.ErrNoProfile
	}

	var profile model.Profile
	err := json.Unmarshal([]byte(events[0].Content), &profile)
	if err != nil {
		return nil, err
	}

	npub, err := publicKey.ToNPublicKey()
	if err != nil {
		return nil, err
	}

	return &usecase.ProfileGetterOutput{
		Profile:    profile,
		NpublicKey: npub,
	}, nil
}
