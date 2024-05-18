// Package config read and write configuration files.
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/nbd-wtf/go-nostr/nip19"
)

const (
	// appName is the name of the application.
	appName = "honeycomb"
	// privateKeyFilePath is the path to the private key file.
	privateKeyFilePath = "private_key"
)

// PrivateKey is the type of the private key.
type PrivateKey string

// String returns the string representation of the private key.
func (p PrivateKey) String() string {
	return string(p)
}

// Validate returns an error if the private key is empty.
func (p PrivateKey) Validate() error {
	if p == "" {
		return ErrEmptyPrivateKey
	}
	if _, _, err := nip19.Decode(p.String()); err != nil {
		return fmt.Errorf("%w: input=%s: %s", ErrInvalidPrivateKey, p.String(), err.Error())
	}
	return nil
}

// DirPath return directory path that store configuration-file.
// Default path is $HOME/.config/honeycomb.
func DirPath() string {
	return filepath.Join(xdg.ConfigHome, appName)
}

// PrivateKeyFilePath return file path that store private key.
// Default path is $HOME/.config/honeycomb/private_key.
func PrivateKeyFilePath() string {
	return filepath.Join(DirPath(), privateKeyFilePath)
}

// ReadPrivateKey reads the private key from the file.
// This function returns an error if the file does not exist or the private key is invalid.
func ReadPrivateKey() (PrivateKey, error) {
	content, err := os.ReadFile(PrivateKeyFilePath())
	if err != nil {
		return "", err
	}
	pk := PrivateKey(content)
	if err := pk.Validate(); err != nil {
		return "", err
	}
	return pk, nil
}

// WritePrivateKey writes the private key to the file.
func WritePrivateKey(pk PrivateKey) error {
	if err := os.MkdirAll(DirPath(), 0700); err != nil {
		return err
	}
	if err := os.WriteFile(PrivateKeyFilePath(), []byte(pk), 0600); err != nil {
		return err
	}
	return nil
}
