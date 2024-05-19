//go:build wireinject
// +build wireinject

// Package di Inject dependence by wire command.
package di

import (
	"context"

	"github.com/google/wire"
	"github.com/nao1215/honeycomb/app/external"
	"github.com/nao1215/honeycomb/app/interactor"
	"github.com/nao1215/honeycomb/app/service"
	"github.com/nao1215/honeycomb/app/usecase"
)

//go:generate go run -mod=mod github.com/google/wire/cmd/wire

// HoneyComb has business logic for honeycomb application.
type HoneyComb struct {
	usecase.ProfileGetter // ProfileGetter is the interface that wraps the basic GetProfile method.
	service.RelayFinder   // RelayFinder is the interface that wraps the basic RelayFinder method.
}

// newHoneyComb creates a new HoneyComb.
func newHoneyComb(
	profileGetter usecase.ProfileGetter,
	relayFinder service.RelayFinder,
) *HoneyComb {
	return &HoneyComb{
		ProfileGetter: profileGetter,
		RelayFinder:   relayFinder,
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
