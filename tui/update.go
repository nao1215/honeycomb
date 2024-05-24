package tui

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/nao1215/honeycomb/app/model"
	"github.com/nao1215/honeycomb/app/usecase"
	"github.com/rivo/tview"
)

// updateViews changes the current view.
func (t *TUI) updateViews() error {
	t.clearAllTextViews()
	if err := t.updateViewText(); err != nil {
		return err
	}
	return nil
}

// clearAllTextViews clears the text views.
func (t *TUI) clearAllTextViews() {
	t.timeline.Clear()
	t.trend.Clear()
	t.follow.Clear()
	t.follower.Clear()
	t.profile.Clear()
	t.setting.Clear()
	t.main.Clear()
	t.footer.Clear()
}

// updateViewText sets the text view text.
func (t *TUI) updateViewText() error {
	t.updateHeaderView()
	if err := t.updateMainTextView(); err != nil {
		return err
	}
	t.updateFooter()
	return nil
}

// updateMainTextView sets the main text view text.
func (t *TUI) updateMainTextView() error {
	switch *t.viewModel.currentView {
	case currentViewTimeline:
		if err := t.writeTimeline(); err != nil {
			return err
		}
	case currentViewTrend:
		t.main.SetText("Trend: not implemented yet")
	case currentViewFollow:
		if err := t.writeFollows(); err != nil {
			return err
		}
	case currentViewFollower:
		t.main.SetText("Follower: not implemented yet")
	case currentViewProfile:
		if err := t.writeProfile(); err != nil {
			return err
		}
	case currentViewSetting:
		t.main.SetText("Setting: not implemented yet")
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

	if *t.viewModel.currentView == currentViewTrend {
		t.trend.SetText(t.viewModel.currentView.stringWithBee())
	} else {
		t.trend.SetText(currentViewTrend.string())
	}

	if *t.viewModel.currentView == currentViewFollow {
		t.follow.SetText(t.viewModel.currentView.stringWithBee())
	} else {
		t.follow.SetText(currentViewFollow.string())
	}

	if *t.viewModel.currentView == currentViewFollower {
		t.follower.SetText(t.viewModel.currentView.stringWithBee())
	} else {
		t.follower.SetText(currentViewFollower.string())
	}

	if *t.viewModel.currentView == currentViewProfile {
		t.profile.SetText(t.viewModel.currentView.stringWithBee())
	} else {
		t.profile.SetText(currentViewProfile.string())
	}

	if *t.viewModel.currentView == currentViewSetting {
		t.setting.SetText(t.viewModel.currentView.stringWithBee())
	} else {
		t.setting.SetText(currentViewSetting.string())
	}
}

// updateFooter sets the footer text.
func (t *TUI) updateFooter() {
	switch *t.viewModel.currentView {
	case currentViewTimeline:
		t.footer.SetText("  Quit:<ESC>, 'q' | Next:<TAB> | Prev:<SHIFT-TAB> | Post:'p'")
	case currentViewTrend:
		t.footer.SetText("  Quit:<ESC>, 'q' | Next:<TAB> | Prev:<SHIFT-TAB> | Post:'p'")
	case currentViewFollow:
		t.footer.SetText("  Quit:<ESC>, 'q' | Next:<TAB> | Prev:<SHIFT-TAB> | Post:'p'")
	case currentViewFollower:
		t.footer.SetText("  Quit:<ESC>, 'q' | Next:<TAB> | Prev:<SHIFT-TAB> | Post:'p'")
	case currentViewProfile:
		t.footer.SetText("  Quit:<ESC>, 'q' | Next:<TAB> | Prev:<SHIFT-TAB> | Post:'p'")
	case currentViewSetting:
		t.footer.SetText("  Quit:<ESC>, 'q' | Next:<TAB> | Prev:<SHIFT-TAB> | Post:'p'")
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

	width, _, err := screenSize()
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

// screenSize returns the screen size.
func screenSize() (int, int, error) {
	s, _ := tcell.NewScreen()
	if err := s.Init(); err != nil {
		return 0, 0, err
	}
	cols, rows := s.Size()
	s.Fini()
	return cols, rows, nil
}

// countWrappedLines counts the number of lines that text will occupy when wrapped to the given width.
func countWrappedLines(text string, width int) int {
	lines := strings.Split(text, "\n")
	fmt.Printf("%s\nlines=%d, width=%d", text, len(lines), width)
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
		AddButtons([]string{"Reply", "Repost", "Like", "Unlike", "Zap"}).
		SetDoneFunc(func(buttonIndex int, _ string) {
			switch buttonIndex {
			case 0: // Reply
				// TODO: implement
			case 1: // Repost
				// TODO: implement
			case 2: // Like
				// TODO: implement
			case 3: // Unlike
				// TODO:implement
			case 4: // Zap
				// TODO: implement
			}
			t.postModalVisible.invisible()
			t.app.SetRoot(t.verticalFlex, true)
		})

	t.postModalVisible.visible()
	t.app.SetRoot(modal, true)
}
