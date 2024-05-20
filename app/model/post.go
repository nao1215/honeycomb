package model

// Post is the user's post.
type Post struct {
	ID        string    // ID is the post(event) ID.
	Author    Profile   // Author is the post author.
	PublicKey PublicKey // PublicKey is the author's public key.
	Content   string    // Content is the post content.
}
