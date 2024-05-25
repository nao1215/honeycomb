package external

import (
	"github.com/google/wire"
	"github.com/nao1215/honeycomb/app/service"
)

// Set is a provider set for external services.
//
//nolint:gochecknoglobals
var Set = wire.NewSet(
	NewRelayFinder,
	wire.Bind(new(service.RelayFinder), new(*RelayFinder)),
	NewEventsLister,
	wire.Bind(new(service.EventsLister), new(*EventsLister)),
	NewPublisher,
	wire.Bind(new(service.Publisher), new(*Publisher)),
)
