package interactor

import (
	"context"

	"github.com/nao1215/honeycomb/app/service"
	"github.com/nao1215/honeycomb/app/usecase"
	"github.com/nbd-wtf/go-nostr"
)

var _ usecase.Liker = (*Liker)(nil)

// Liker implements the Liker interface.
type Liker struct {
	service.Publisher
}

// NewLiker creates a new Liker.
func NewLiker(
	publisher service.Publisher,
) *Liker {
	return &Liker{
		Publisher: publisher,
	}
}

// Like likes a message.
func (l *Liker) Like(ctx context.Context, input *usecase.LikerInput) (*usecase.LikerOutput, error) {
	event := nostr.Event{
		Kind:      nostr.KindReaction,
		PubKey:    input.NPublicKey.String(),
		Tags:      nostr.Tags{nostr.Tag{"e", input.PostID}},
		CreatedAt: nostr.Now(),
		Content:   "+",
	}

	if err := event.Sign(input.PrivateKey.String()); err != nil {
		return nil, err
	}

	if _, err := l.Publisher.Publish(
		ctx,
		&service.PublisherInput{
			Event:          event,
			ConnectedRelay: input.ConnectedRelay,
		}); err != nil {
		return nil, err
	}
	return &usecase.LikerOutput{}, nil
}
