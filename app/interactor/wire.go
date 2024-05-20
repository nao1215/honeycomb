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
	NewAuthorGetter,
	wire.Bind(new(usecase.AuthorGetter), new(*AuthorGetter)),
	NewFollowLister,
	wire.Bind(new(usecase.FollowLister), new(*FollowLister)),
	NewPoster,
	wire.Bind(new(usecase.Poster), new(*Poster)),
	NewTimelineLister,
	wire.Bind(new(usecase.TimelineLister), new(*TimelineLister)),
)
