package service

import (
	"context"

	"github.com/nbd-wtf/go-nostr"
)

// PublisherInput is the input for the publishing event
type PublisherInput struct {
	Event          nostr.Event  // Event is the nostr event.
	ConnectedRelay *nostr.Relay // ConnectedRelay is the connected relay.
}

// PublisherOutput is the output for the publishing event
type PublisherOutput struct{}

// Publisher is the interface that wraps the basic Publish method.
type Publisher interface {
	Publish(ctx context.Context, input *PublisherInput) (*PublisherOutput, error)
}
