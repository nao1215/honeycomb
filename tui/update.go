package tui

import (
	"fmt"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/nao1215/honeycomb/app/model"
	"github.com/nao1215/honeycomb/app/usecase"
	"github.com/rivo/tview"
	"github.com/shogo82148/pointer"
	"golang.org/x/term"
)

// updateViews changes the current view.
func (t *TUI) updateViews() error {
	if err := t.updateViewText(); err != nil {
		return err
	}
	return nil
}

// updateViewText sets the text view text.
func (t *TUI) updateViewText() error {
	t.updateHeaderView()
	t.updateFooter()
	return t.updateMainTextView()
}

// updateMainTextView sets the main text view text.
func (t *TUI) updateMainTextView() error {
	t.main.Clear()
	switch *t.viewModel.currentView {
	case currentViewTimeline:
		if err := t.writeTimeline(); err != nil {
			return err
		}
		t.main.ScrollTo(0, 0)
	case currentViewFollow:
		if err := t.writeFollows(); err != nil {
			return err
		}
		t.main.ScrollTo(0, 0)
	case currentViewProfile:
		if err := t.writeProfile(); err != nil {
			return err
		}
		t.main.ScrollTo(0, 0)
	}
	return nil
}

// updateHeaderView highlights the header.
func (t *TUI) updateHeaderView() {
	if *t.viewModel.currentView == currentViewTimeline {
		t.timeline.SetText(t.viewModel.currentView.stringWithBee())
	} else {
		t.timeline.SetText(currentViewTimeline.string())
	}

	if *t.viewModel.currentView == currentViewFollow {
		t.follow.SetText(t.viewModel.currentView.stringWithBee())
	} else {
		t.follow.SetText(currentViewFollow.string())
	}

	if *t.viewModel.currentView == currentViewProfile {
		t.profile.SetText(t.viewModel.currentView.stringWithBee())
	} else {
		t.profile.SetText(currentViewProfile.string())
	}
}

// updateFooter sets the footer text.
func (t *TUI) updateFooter() {
	switch *t.viewModel.currentView {
	case currentViewTimeline:
		t.footer.SetText("  Quit:<ESC> | Next:<TAB> | Prev:<SHIFT-TAB> | Post:'p' | Reload: 'R' | Reaction: <Enter>")
	case currentViewFollow:
		t.footer.SetText("  Quit:<ESC> | Next:<TAB> | Prev:<SHIFT-TAB> | Post:'p' | Reload: 'R'")
	case currentViewProfile:
		t.footer.SetText("  Quit:<ESC> | Next:<TAB> | Prev:<SHIFT-TAB> | Post:'p'")
	}
}

const (
	// headerOffsetLineCount is the number of lines occupied by the header.
	headerOffsetLineCount = 4
)

// writeTimeline writes timeline posts to the main text view.
// It also sets the post ranges. The post ranges are used to determine which post is selected.
func (t *TUI) writeTimeline() error {
	lineCount := headerOffsetLineCount
	t.viewModel.postRanges = nil // Clear previous post ranges

	width, _, err := terminalSize()
	if err != nil {
		return err
	}
	for _, post := range t.viewModel.timeline {
		postText := fmt.Sprintf("[yellow]%s[white]\n%s\n\n", post.Author.DisplayNameOrName(), post.Content)
		totalLines := countWrappedLines(fmt.Sprintf("%s\n%s\n\n", post.Author.DisplayNameOrName(), post.Content), width)
		postRange := postRange{
			post:      post,
			startLine: lineCount,
			lineCount: totalLines,
		}
		t.viewModel.postRanges = append(t.viewModel.postRanges, &postRange)
		lineCount += totalLines + 1 // +1 for the extra newline after the post

		if _, err := t.main.Write([]byte(postText)); err != nil {
			return err
		}
	}
	return nil
}

func terminalSize() (width int, height int, err error) {
	fd := int(os.Stdin.Fd())
	width, height, err = term.GetSize(fd)
	if err != nil {
		return 0, 0, err
	}
	return width, height, nil
}

// countWrappedLines counts the number of lines that text will occupy when wrapped to the given width.
func countWrappedLines(text string, width int) int {
	lines := strings.Split(text, "\n")
	lineCount := 0

	for _, line := range lines {
		words := strings.Fields(line)
		currentLineLength := 0

		for _, word := range words {
			if currentLineLength+len(word)+1 > width { // +1 for the space
				lineCount++
				currentLineLength = 0
			}
			currentLineLength += len(word) + 1
		}

		if currentLineLength > 0 {
			lineCount++
		}
	}
	return lineCount
}

// writeFollows writes follows to the main text view.
func (t *TUI) writeFollows() error {
	for _, follow := range *t.viewModel.follows {
		displayName := follow.Profile.DisplayNameOrName()

		website := follow.Profile.Website
		if website == "" {
			website = "No website"
		}

		text := fmt.Sprintf("[yellow]%s[white]\nWebsite=%s\n%s\n\n",
			displayName,
			website,
			follow.Profile.About)

		if _, err := t.main.Write([]byte(text)); err != nil {
			return err
		}
	}
	return nil
}

// writeProfile writes the profile to the main text view.
func (t *TUI) writeProfile() error {
	profile := t.viewModel.author.Profile

	displayName := profile.DisplayNameOrName()
	website := profile.Website
	if website == "" {
		website = "No website"
	}

	text := fmt.Sprintf("[yellow]%s[white]\nWebsite=%s\n%s\n\n\n\n",
		displayName,
		website,
		profile.About)
	if _, err := t.main.Write([]byte(text)); err != nil {
		return err
	}

	for _, post := range t.viewModel.myPosts {
		displayName := post.Author.DisplayNameOrName()
		postText := fmt.Sprintf("[yellow]%s[white]\n%s\n\n", displayName, post.Content)

		if _, err := t.main.Write([]byte(postText)); err != nil {
			return err
		}
	}
	return nil
}

// writePost handles the post action.
func (t *TUI) writePost() {
	textArea, ok := t.postForm.GetFormItem(0).(*tview.TextArea)
	if !ok {
		showError(t.app, "Failed to read post input field.")
		return
	}
	text := textArea.GetText()
	if text == "" {
		showError(t.app, "Post content is empty.")
		return
	}

	_, err := t.honeycomb.Post(t.ctx, &usecase.PosterInput{
		Content:        text,
		PrivateKey:     t.viewModel.author.PrivateKey,
		NPublicKey:     t.viewModel.author.NPublicKey,
		ConnectedRelay: t.viewModel.author.ConnectedRelay,
	})
	if err != nil {
		showError(t.app, err.Error())
		return
	}
	t.postFormVisible.invisible()
	t.app.SetRoot(t.verticalFlex, true)
}

// showPostModal displays a modal with the post content and buttons.
func (t *TUI) showPostModal(post *model.Post) {
	postText := fmt.Sprintf("[yellow]%s\n[white]\n%s", post.Author.DisplayNameOrName(), post.Content)
	modal := tview.NewModal().
		SetText(postText).
		AddButtons([]string{"Like"}).
		SetDoneFunc(func(buttonIndex int, _ string) {
			switch buttonIndex {
			case 0: // Like
				if _, err := t.honeycomb.Like(t.ctx, &usecase.LikerInput{
					PostID:         post.ID,
					PrivateKey:     t.viewModel.author.PrivateKey,
					ConnectedRelay: t.viewModel.author.ConnectedRelay,
				}); err != nil {
					t.postModalVisible.invisible()
					t.app.SetRoot(t.verticalFlex, true)
					showError(t.app, err.Error())
					return
				}
			}
			t.postModalVisible.invisible()
			t.app.SetRoot(t.verticalFlex, true)
		})

	modal.SetBackgroundColor(tcell.ColorDefault)
	modal.SetButtonStyle(tcell.StyleDefault)
	modal.SetBorder(false)

	t.postModalVisible.visible()
	t.app.SetRoot(modal, true)
}

// appendOldTimeline appends the old timeline to the current timeline.
func (t *TUI) appendOldTimeline() error {
	oldestPost := t.viewModel.timeline[len(t.viewModel.timeline)-1]

	output, err := t.honeycomb.TimelineLister.ListTimeline(t.ctx, &usecase.TimelineListerInput{
		Follows:        pointer.Value(t.viewModel.follows),
		Until:          pointer.Ptr(oldestPost.CreatedAt),
		ConnectedRelay: t.viewModel.author.ConnectedRelay,
	})
	if err != nil {
		return err
	}
	t.viewModel.timeline = append(t.viewModel.timeline, output.Posts...)
	t.viewModel.timeline.ToUniquePosts()
	return nil
}

// appendOldTimelineAndRewriteIfNeeded checks if the timeline needs to be appended and appends if needed.
func (t *TUI) appendOldTimelineAndRewriteIfNeeded() error {
	threshold := len(t.viewModel.timeline) - 12

	row, _ := t.main.GetScrollOffset()
	for i, v := range t.viewModel.postRanges {
		if row >= v.startLine {
			if i >= threshold {
				if err := t.appendOldTimeline(); err != nil {
					return err
				}
				return t.rewriteTimeline()
			}
		}
	}
	return nil
}

// rewriteTimeline redraws the timeline while preserving the scroll position.
func (t *TUI) rewriteTimeline() error {
	row, column := t.main.GetScrollOffset()
	if err := t.updateViewText(); err != nil {
		return err
	}
	t.main.ScrollTo(row, column)
	return nil
}
