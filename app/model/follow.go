package model

// Follow is the user's follower.
type Follow struct {
	PublicKey PublicKey // PublicKey is the user's public key.
	Profile   Profile   // Profile is the user's profile.
}
