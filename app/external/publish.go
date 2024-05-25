package external

import (
	"context"
	"fmt"

	"github.com/nao1215/honeycomb/app/service"
)

var _ service.Publisher = (*Publisher)(nil)

// Publisher is the external service for posting a message.
type Publisher struct{}

// NewPublisher creates a new Publisher.
func NewPublisher() *Publisher {
	return &Publisher{}
}

// Publish publish a message.
func (p *Publisher) Publish(ctx context.Context, input *service.PublisherInput) (*service.PublisherOutput, error) {
	err := input.ConnectedRelay.Publish(ctx, input.Event)
	if err != nil {
		return nil, fmt.Errorf("failed to publish the event: %w", err)
	}
	return &service.PublisherOutput{}, nil
}
