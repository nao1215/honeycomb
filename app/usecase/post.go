package usecase

import (
	"context"

	"github.com/nao1215/honeycomb/app/model"
	"github.com/nbd-wtf/go-nostr"
)

// PosterInput is the input of the Post method.
type PosterInput struct {
	Content        string // Content is the post content.
	PrivateKey     model.PrivateKey
	NPublicKey     model.NPublicKey // NPublicKey is the user's public key.
	ConnectedRelay *nostr.Relay     // ConnectedRelay is the connected relay.
}

// PosterOutput is the output of the Post method.
type PosterOutput struct{}

// Poster is the interface that wraps the basic Poster method.
type Poster interface {
	Post(ctx context.Context, input *PosterInput) (*PosterOutput, error)
}
