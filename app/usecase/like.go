package usecase

import (
	"context"

	"github.com/nao1215/honeycomb/app/model"
	"github.com/nbd-wtf/go-nostr"
)

// LikerInput is the input of the Like method.
type LikerInput struct {
	PostID         string // PostID is the post id.
	PrivateKey     model.PrivateKey
	NPublicKey     model.NPublicKey
	ConnectedRelay *nostr.Relay // ConnectedRelay is the connected relay.
}

// LikerOutput is the output of the Like method.
type LikerOutput struct{}

// Liker is the interface that wraps the basic Like method.
type Liker interface {
	Like(ctx context.Context, input *LikerInput) (*LikerOutput, error)
}
