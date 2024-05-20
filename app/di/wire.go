//go:build wireinject
// +build wireinject

// Package di Inject dependence by wire command.
package di

import (
	"context"

	"github.com/google/wire"
	"github.com/nao1215/honeycomb/app/external"
	"github.com/nao1215/honeycomb/app/interactor"
	"github.com/nao1215/honeycomb/app/usecase"
)

//go:generate go run -mod=mod github.com/google/wire/cmd/wire

// HoneyComb has business logic for honeycomb application.
type HoneyComb struct {
	usecase.AuthorGetter   // AuthorGetter is the interface that wraps the basic GetAuthor method.
	usecase.ProfileGetter  // ProfileGetter is the interface that wraps the basic GetProfile method.
	usecase.FollowLister   // FollowLister is the interface that wraps the basic ListFollow method.
	usecase.Poster         // Poster is the interface that wraps the basic Post method.
	usecase.TimelineLister // TimelineLister is the interface that wraps the basic ListTimeline method.
}

// newHoneyComb creates a new HoneyComb.
func newHoneyComb(
	authorGetter usecase.AuthorGetter,
	profileGetter usecase.ProfileGetter,
	followLister usecase.FollowLister,
	poster usecase.Poster,
	timelineLister usecase.TimelineLister,
) *HoneyComb {
	return &HoneyComb{
		AuthorGetter:   authorGetter,
		ProfileGetter:  profileGetter,
		FollowLister:   followLister,
		Poster:         poster,
		TimelineLister: timelineLister,
	}
}

// NewHoneyComb creates a new HoneyComb.
func NewHoneyComb(ctx context.Context) (*HoneyComb, error) {
	wire.Build(
		interactor.Set,
		external.Set,
		newHoneyComb,
	)
	return nil, nil
}
