// Package usecase has interfaces that wrap the basic business logic.
package usecase

import (
	"context"

	"github.com/nao1215/honeycomb/app/model"
	"github.com/nbd-wtf/go-nostr"
)

// ProfileGetterInput is the input of the GetProfile method.
type ProfileGetterInput struct {
	PublicKey      model.PublicKey // PublicKey is the user's public key.
	ConnectedRelay *nostr.Relay    // ConnectedRelay is the connected relay
}

// ProfileGetterOutput is the output of the GetProfile method.
type ProfileGetterOutput struct {
	// Profile is the user's profile.
	Profile model.Profile
	// NPublicKey is the user's public key.
	NPublicKey model.NPublicKey
}

// ProfileGetter is the interface that wraps the basic GetProfile method.
type ProfileGetter interface {
	GetProfile(ctx context.Context, input *ProfileGetterInput) (*ProfileGetterOutput, error)
}
