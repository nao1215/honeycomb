// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"context"
	"github.com/nao1215/honeycomb/app/external"
	"github.com/nao1215/honeycomb/app/interactor"
	"github.com/nao1215/honeycomb/app/usecase"
)

// Injectors from wire.go:

// NewHoneyComb creates a new HoneyComb.
func NewHoneyComb(ctx context.Context) (*HoneyComb, error) {
	relayFinder := external.NewRelayFinder()
	eventsLister := external.NewEventsLister()
	authorGetter := interactor.NewAuthorGetter(relayFinder, eventsLister)
	profileGetter := interactor.NewProfileGetter(relayFinder, eventsLister)
	followLister := interactor.NewFollowLister(relayFinder, eventsLister)
	honeyComb := newHoneyComb(authorGetter, profileGetter, followLister)
	return honeyComb, nil
}

// wire.go:

// HoneyComb has business logic for honeycomb application.
type HoneyComb struct {
	usecase.AuthorGetter
	usecase.
		// AuthorGetter is the interface that wraps the basic GetAuthor method.
		ProfileGetter
	usecase.FollowLister

	// ProfileGetter is the interface that wraps the basic GetProfile method.
	// FollowLister is the interface that wraps the basic ListFollow method.
}

// newHoneyComb creates a new HoneyComb.
func newHoneyComb(
	authorGetter usecase.AuthorGetter,
	profileGetter usecase.ProfileGetter,
	followLister usecase.FollowLister,
) *HoneyComb {
	return &HoneyComb{
		AuthorGetter:  authorGetter,
		ProfileGetter: profileGetter,
		FollowLister:  followLister,
	}
}
