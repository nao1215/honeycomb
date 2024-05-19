package interactor

import (
	"context"

	"github.com/nao1215/honeycomb/app/service"
	"github.com/nao1215/honeycomb/app/usecase"
	"github.com/nbd-wtf/go-nostr"
)

var _ usecase.Poster = (*Poster)(nil)

// Poster implements the Poster interface.
type Poster struct {
	service.Poster
}

// NewPoster creates a new Poster.
func NewPoster(
	poster service.Poster,
) *Poster {
	return &Poster{
		Poster: poster,
	}
}

// Post posts a message.
func (p *Poster) Post(ctx context.Context, input *usecase.PosterInput) (*usecase.PosterOutput, error) {
	// TODO: this event is very simple, we need to implement more complex event creation.
	event := nostr.Event{
		PubKey:    input.NPublicKey.String(),
		Content:   input.Content,
		CreatedAt: nostr.Now(),
		Kind:      nostr.KindTextNote,
		Tags:      nostr.Tags{},
	}

	if err := event.Sign(input.PrivateKey.String()); err != nil {
		return nil, err
	}

	if _, err := p.Poster.Post(ctx, &service.PosterInput{
		Event:          event,
		ConnectedRelay: input.ConnectedRelay,
	}); err != nil {
		return nil, err
	}
	return &usecase.PosterOutput{}, nil
}
