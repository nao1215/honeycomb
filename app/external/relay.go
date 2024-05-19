// Package external implements the external service.
package external

import (
	"context"

	"github.com/nao1215/honeycomb/app/service"
	"github.com/nbd-wtf/go-nostr"
)

var _ service.RelayFinder = (*RelayFinder)(nil)

// RelayFinder is the external service for finding a relay.
type RelayFinder struct{}

// NewRelayFinder creates a new RelayFinder.
func NewRelayFinder() *RelayFinder {
	return &RelayFinder{}
}

// FindRelay finds a relay.
func (r *RelayFinder) FindRelay(ctx context.Context, input *service.RelayFinderInput) (*service.RelayFinderOutput, error) {
	for k, v := range input.Relays {
		if !v.Read {
			continue
		}
		relay, err := nostr.RelayConnect(ctx, k.String())
		if err != nil {
			return nil, err
		}
		return &service.RelayFinderOutput{
			Relay: relay,
		}, nil
	}
	return nil, service.ErrNoRelayConnection
}
