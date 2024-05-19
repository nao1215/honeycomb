// Package usecase has interfaces that wrap the basic business logic.
package usecase

import (
	"context"

	"github.com/nao1215/honeycomb/app/model"
)

// ProfileGetterInput is the input of the GetProfile method.
type ProfileGetterInput struct {
	NsecretKey model.NSecretKey // NsecretKey is the user's private key.
}

// ProfileGetterOutput is the output of the GetProfile method.
type ProfileGetterOutput struct {
	// Profile is the user's profile.
	Profile model.Profile
	// NpublicKey is the user's public key.
	NpublicKey model.NPublicKey
}

// ProfileGetter is the interface that wraps the basic GetProfile method.
type ProfileGetter interface {
	GetProfile(ctx context.Context, input *ProfileGetterInput) (*ProfileGetterOutput, error)
}
