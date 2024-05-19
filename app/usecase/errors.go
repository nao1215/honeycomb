package usecase

import "errors"

var (
	// ErrNoProfile is returned when the profile is not found.
	ErrNoProfile = errors.New("no profile")
)
