package model

import "errors"

var (
	// ErrEmptyPrivateKey is returned when the private key is empty.
	ErrEmptyPrivateKey = errors.New("private key is empty")
	// ErrInvalidPrivateKey is returned when the private key is invalid.
	ErrInvalidPrivateKey = errors.New("private key is invalid")
)
