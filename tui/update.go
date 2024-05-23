package tui

import (
	"fmt"

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
		t.main.SetText("Trend")
	case currentViewFollow:
		if err := t.writeFollows(); err != nil {
			return err
		}
	case currentViewFollower:
		t.main.SetText("Follower")
	case currentViewProfile:
		if err := t.writeProfile(); err != nil {
			return err
		}
	case currentViewSetting:
		t.main.SetText("Setting")
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

// writeTimeline writes timeline posts to the main text view.
// It also sets the post ranges. The post ranges are used to determine which post is selected.
func (t *TUI) writeTimeline() error {
	lineCount := 0
	t.viewModel.postRanges = nil // Clear previous post ranges

	for _, post := range t.viewModel.timeline {
		displayName := post.Author.DisplayName
		if displayName == "" {
			displayName = post.Author.Name
		}
		postText := fmt.Sprintf("[yellow]%s[white]\n%s\n\n", displayName, post.Content)
		_, _, width, _ := t.main.GetInnerRect()
		lines := tview.WordWrap(postText, width)

		postRange := postRange{
			post:      post,
			startLine: lineCount,
			lineCount: len(lines),
		}
		t.viewModel.postRanges = append(t.viewModel.postRanges, &postRange)
		lineCount += len(lines)

		if _, err := t.main.Write([]byte(postText)); err != nil {
			return err
		}
	}
	return nil
}

// writeFollows writes follows to the main text view.
func (t *TUI) writeFollows() error {
	for _, follow := range *t.viewModel.follows {
		displayName := follow.Profile.DisplayName
		if displayName == "" {
			displayName = follow.Profile.Name
		}

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

	displayName := profile.DisplayName
	if displayName == "" {
		displayName = profile.Name
	}

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
		displayName := post.Author.DisplayName
		if displayName == "" {
			displayName = post.Author.Name
		}
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
		return // TODO: error handling
	}
	text := textArea.GetText()
	if text == "" {
		return
	}

	_, err := t.honeycomb.Post(t.ctx, &usecase.PosterInput{
		Content:        text,
		PrivateKey:     t.viewModel.author.PrivateKey,
		NPublicKey:     t.viewModel.author.NPublicKey,
		ConnectedRelay: t.viewModel.author.ConnectedRelay,
	})
	if err != nil {
		// TODO: error handling
		panic(err)
	}
	t.postFormVisible.invisible()
	t.app.SetRoot(t.verticalFlex, true)
}
