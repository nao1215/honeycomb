package interactor

import (
	"github.com/google/wire"
	"github.com/nao1215/honeycomb/app/usecase"
)

// Set is the wire set for the interactor package.
//
//nolint:gochecknoglobals
var Set = wire.NewSet(
	NewProfileGetter,
	wire.Bind(new(usecase.ProfileGetter), new(*ProfileGetter)),
)
