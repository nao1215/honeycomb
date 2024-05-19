package model

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/nao1215/honeycomb/config"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

const (
	// nSecretKeyFilePath is the path to the private key file.
	nSecretKeyFilePath = "private_key"
)

// PrivateKey is the type of the private key.
type PrivateKey string

// String returns the string representation of the private key.
func (p PrivateKey) String() string {
	return string(p)
}

// PublicKey is the type of the public key.
type PublicKey string

// String returns the string representation of the public key.
func (p PublicKey) String() string {
	return string(p)
}

// ToNPublicKey returns the nostr public key.
func (p PublicKey) ToNPublicKey() (NPublicKey, error) {
	npub, err := nip19.EncodePublicKey(p.String())
	if err != nil {
		return "", err
	}
	return NPublicKey(npub), nil
}

// NPublicKey is the type of the nostr public key.
type NPublicKey string

// String returns the string representation of the public key.
func (p NPublicKey) String() string {
	return string(p)
}

// NSecretKey is the type of the private key.
type NSecretKey string

// String returns the string representation of the private key.
func (p NSecretKey) String() string {
	return string(p)
}

// Validate returns an error if the private key is empty.
func (p NSecretKey) Validate() error {
	if p == "" {
		return ErrEmptyPrivateKey
	}
	if _, _, err := nip19.Decode(p.String()); err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidPrivateKey, err.Error())
	}
	return nil
}

// ToPrivateKey returns the private key from the private key.
func (p NSecretKey) ToPrivateKey() (PrivateKey, error) {
	_, pk, err := nip19.Decode(p.String())
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrInvalidPrivateKey, err.Error())
	}

	pkStr, ok := pk.(string)
	if !ok {
		return "", fmt.Errorf("failed to convert private key to string")
	}
	return PrivateKey(pkStr), nil
}

// ToPublicKey returns the public key from the private key.
func (p NSecretKey) ToPublicKey() (PublicKey, error) {
	_, s, err := nip19.Decode(p.String())
	if err != nil {
		return "", err
	}

	privateKey, ok := s.(string)
	if !ok {
		return "", fmt.Errorf("failed to convert private key to string")
	}

	pub, err := nostr.GetPublicKey(privateKey)
	if err != nil {
		return "", err
	}
	return PublicKey(pub), nil
}

// DirPath return directory path that store configuration-file.
// Default path is $HOME/.config/honeycomb.
func DirPath() string {
	return filepath.Join(xdg.ConfigHome, config.AppName)
}

// NSecretKeyFilePath return file path that store private key.
// Default path is $HOME/.config/honeycomb/private_key.
func NSecretKeyFilePath() string {
	return filepath.Join(DirPath(), nSecretKeyFilePath)
}

// ReadNSecretKey reads the private key from the file.
// This function returns an error if the file does not exist or the private key is invalid.
func ReadNSecretKey() (NSecretKey, error) {
	content, err := os.ReadFile(NSecretKeyFilePath())
	if err != nil {
		return "", err
	}
	pk := NSecretKey(content)
	if err := pk.Validate(); err != nil {
		return "", err
	}
	return pk, nil
}

// WriteNSecretKey writes the private key to the file.
func WriteNSecretKey(pk NSecretKey) error {
	if err := os.MkdirAll(DirPath(), 0700); err != nil {
		return err
	}
	if err := os.WriteFile(NSecretKeyFilePath(), []byte(pk), 0600); err != nil {
		return err
	}
	return nil
}
