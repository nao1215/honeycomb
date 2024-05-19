package usecase

import (
	"context"

	"github.com/nao1215/honeycomb/app/model"
	"github.com/nbd-wtf/go-nostr"
)

// FollowListerInput is the input of the FollowersLister method.
type FollowListerInput struct {
	PublicKey      model.PublicKey // PublicKey is the user's public key.
	ConnectedRelay *nostr.Relay    // ConnectedRelay is the connected relay
}

// FollowListerOutput is the output of the FollowLister method.
type FollowListerOutput struct {
	Follows []*model.Follow // Follows is the list of follow user.
}

// FollowLister is the interface that wraps the basic FollowLister method.
type FollowLister interface {
	ListFollow(ctx context.Context, input *FollowListerInput) (*FollowListerOutput, error)
}
