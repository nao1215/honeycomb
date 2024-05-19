package service

import (
	"context"

	"github.com/nbd-wtf/go-nostr"
)

// EventsListerInput is the input for the EventsLister.
type EventsListerInput struct {
	Filter nostr.Filter // Filter is the filter.
	// TODO: Use multiple relays.
	Relay *nostr.Relay // Relay is the relay. Already connected.
}

// EventsListerOutput is the output for the EventsLister.
type EventsListerOutput struct {
	Events []*nostr.Event // Events is the list of events.
}

// EventsLister is the interface that wraps the basic ListEvents method.
type EventsLister interface {
	ListEvents(ctx context.Context, input *EventsListerInput) (*EventsListerOutput, error)
}
