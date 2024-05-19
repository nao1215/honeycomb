// Package service implements the service layer. This package is used only for application domain.
package service

import (
	"context"

	"github.com/nao1215/honeycomb/app/model"
	"github.com/nbd-wtf/go-nostr"
)

// RelayFinderInput is the input of the RelayFinder method.
type RelayFinderInput struct {
	Relays map[model.WSS]model.Relay // Relays is the list of relays.
}

// RelayFinderOutput is the output of the RelayFinder method.
type RelayFinderOutput struct {
	// TODO: Change output to honeycomb defined type.
	Relay *nostr.Relay // Relay is the relay. If not found, it is nil. You must close it after use.
}

// RelayFinder is the interface that wraps the basic RelayFinder method.
type RelayFinder interface {
	FindRelay(ctx context.Context, input *RelayFinderInput) (*RelayFinderOutput, error)
}
