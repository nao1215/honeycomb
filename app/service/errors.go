package service

import "errors"

var (
	// ErrNoRelayConnection is returned when there is no connection to the relay.
	ErrNoRelayConnection = errors.New("no relay connection")
)
