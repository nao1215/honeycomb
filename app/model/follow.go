package model

// Follow is the user's follower.
type Follow struct {
	PublicKey PublicKey // PublicKey is the user's public key.
	Profile   Profile   // Profile is the user's profile.
}

// Follows is the list of followers.
type Follows []*Follow

// PublicKeys returns the list of public keys.
func (f Follows) PublicKeys() []PublicKey {
	pks := make([]PublicKey, len(f))
	for i, follow := range f {
		pks[i] = follow.PublicKey
	}
	return pks
}

// PublicKeyToFollowMap converts the list of followers to the map of public key to follow.
func (f Follows) PublicKeyToFollowMap() map[PublicKey]*Follow {
	m := make(map[PublicKey]*Follow, len(f))
	for _, follow := range f {
		m[follow.PublicKey] = follow
	}
	return m
}
