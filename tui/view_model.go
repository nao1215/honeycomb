package tui

import "github.com/nao1215/honeycomb/app/model"

// viewModel is the view model.
type viewModel struct {
	author      *model.Author
	follows     *model.Follows
	timeline    []*model.Post
	myPosts     []*model.Post
	currentView *currentView
	postRanges  []*postRange
}

// Close closes the view model.
func (vm *viewModel) Close() error {
	return vm.author.Close()
}

// postRange represents the range of lines occupied by a post in the TextView.
type postRange struct {
	post      *model.Post // Post is the post.
	startLine int         // StartLine is the start line of the post in the TextView.
	lineCount int         // LineCount is the number of lines occupied by the post in the TextView.
}
