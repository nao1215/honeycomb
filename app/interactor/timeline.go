package interactor

import (
	"context"

	"github.com/nao1215/honeycomb/app/model"
	"github.com/nao1215/honeycomb/app/service"
	"github.com/nao1215/honeycomb/app/usecase"
)

var _ usecase.TimelineLister = (*TimelineLister)(nil)

// TimelineLister implements the TimelineLister interface.
type TimelineLister struct {
	service.EventsLister
}

// NewTimelineLister creates a new TimelineLister.
func NewTimelineLister(
	timelineLister service.EventsLister,
) *TimelineLister {
	return &TimelineLister{
		EventsLister: timelineLister,
	}
}

// ListTimeline gets the user's timeline.
func (t *TimelineLister) ListTimeline(ctx context.Context, input *usecase.TimelineListerInput) (*usecase.TimelineListerOutput, error) {
	eventsOutput, err := t.EventsLister.ListEvents(ctx, &service.EventsListerInput{
		Filter: model.FollowsTimelineFilter(
			input.Follows.PublicKeys(),
			input.Since,
			input.Until,
			input.Limit,
		),
		Relay: input.ConnectedRelay,
	})
	if err != nil {
		return nil, err
	}

	posts := make([]*model.Post, 0, len(eventsOutput.Events))
	publicKeyToFollow := input.Follows.PublicKeyToFollowMap()

	for _, event := range eventsOutput.Events {
		post := &model.Post{}
		follow, ok := publicKeyToFollow[model.PublicKey(event.PubKey)]
		if !ok {
			continue
		}

		post.Author = follow.Profile
		post.Content = event.Content
		post.CreatedAt = event.CreatedAt
		posts = append(posts, post)
	}
	return &usecase.TimelineListerOutput{
		Posts: posts,
	}, nil
}
