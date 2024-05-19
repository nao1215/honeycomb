package usecase

import (
	"context"

	"github.com/nao1215/honeycomb/app/model"
)

// AuthorGetterInput is the input of the GetAuthor method.
type AuthorGetterInput struct {
	NSecretKey model.NSecretKey // NSecretKey is the user's private key.
}

// AuthorGetterOutput is the output of the GetAuthor method.
type AuthorGetterOutput struct {
	// Author is the your(honeycomb user) profile data.
	Author *model.Author
}

// AuthorGetter is the interface that wraps the basic GetAuthor method.
type AuthorGetter interface {
	GetAuthor(ctx context.Context, input *AuthorGetterInput) (*AuthorGetterOutput, error)
}
