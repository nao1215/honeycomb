package model

import (
	"github.com/nbd-wtf/go-nostr"
)

// Post is the user's post.
type Post struct {
	ID        string          // ID is the post(event) ID.
	Author    Profile         // Author is the post author.
	PublicKey PublicKey       // PublicKey is the author's public key.
	Content   string          // Content is the post content.
	CreatedAt nostr.Timestamp // CreatedAt is the post creation time.
}

// Posts is the list of posts.
type Posts []*Post

// ToUniquePosts converts the list of posts to the unique list of posts.
func (p Posts) ToUniquePosts() {
	unique := make(map[string]*Post, len(p))
	for _, post := range p {
		unique[post.ID] = post
	}

	p = make([]*Post, 0, len(unique))
	for _, post := range unique {
		p = append(p, post)
	}
}
