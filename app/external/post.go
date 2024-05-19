package external

import (
	"context"
	"fmt"

	"github.com/nao1215/honeycomb/app/service"
)

var _ service.Poster = (*Poster)(nil)

// Poster is the external service for posting a message.
type Poster struct{}

// NewPoster creates a new Poster.
func NewPoster() *Poster {
	return &Poster{}
}

// Post posts a message.
func (p *Poster) Post(ctx context.Context, input *service.PosterInput) (*service.PosterOutput, error) {
	err := input.ConnectedRelay.Publish(ctx, input.Event)
	if err != nil {
		return nil, fmt.Errorf("failed to publish the event: %w", err)
	}
	return &service.PosterOutput{}, nil
}
