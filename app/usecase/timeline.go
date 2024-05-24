package usecase

import (
	"context"

	"github.com/nao1215/honeycomb/app/model"
	"github.com/nbd-wtf/go-nostr"
)

// TimelineListerInput is the input of the ListTimeline method.
type TimelineListerInput struct {
	Follows        model.Follows    // Follows is the list of follow user.
	Since          *nostr.Timestamp // Since is the time to get posts from.
	Until          *nostr.Timestamp // Until is the time to get posts until.
	Limit          int              // Limit is the number of posts to get.
	ConnectedRelay *nostr.Relay     // ConnectedRelay is the connected relay
}

// TimelineListerOutput is the output of the ListTimeline method.
type TimelineListerOutput struct {
	Posts []*model.Post // Posts is the list of posts. The list is sorted by the time of the post.
}

// TimelineLister is the interface that wraps the basic ListTimeline method.
type TimelineLister interface {
	ListTimeline(ctx context.Context, input *TimelineListerInput) (*TimelineListerOutput, error)
}
