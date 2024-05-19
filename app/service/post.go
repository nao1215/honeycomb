package service

import (
	"context"

	"github.com/nbd-wtf/go-nostr"
)

// PosterInput is the input for the Post.
type PosterInput struct {
	Event          nostr.Event  // Event is the nostr event.
	ConnectedRelay *nostr.Relay // ConnectedRelay is the connected relay.
}

// PosterOutput is the output for the Post.
type PosterOutput struct{}

// Poster is the interface that wraps the basic Post method.
type Poster interface {
	Post(ctx context.Context, input *PosterInput) (*PosterOutput, error)
}
