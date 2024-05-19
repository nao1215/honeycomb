package model

import "github.com/nbd-wtf/go-nostr"

// Author is the nostr user, it's you.
type Author struct {
	NSecretKey     NSecretKey    // NSecretKey is the user's nsec private key.
	NPublicKey     NPublicKey    // NPublicKey is the user's npub public key.
	PublicKey      PublicKey     // PublicKey is the user's public key.
	PrivateKey     PrivateKey    // PrivateKey is the user's private key.
	Profile        *Profile      // Profile is the user's profile.
	Relays         map[WSS]Relay // Relays is the list of relays that the user has.
	ConnectedRelay *nostr.Relay  // ConnectedRelay is the connected relay. If not found, it is nil. You must close it after use.
}

// Close closes the connected relay.
func (a *Author) Close() error {
	if a.ConnectedRelay == nil {
		return nil
	}
	if err := a.ConnectedRelay.Close(); err != nil {
		return err
	}
	return nil
}
