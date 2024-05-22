package tui

import (
	"context"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/nao1215/honeycomb/app/di"
	"github.com/nao1215/honeycomb/app/model"
	"github.com/rivo/tview"
)

// TUI is the text-based user interface.
type TUI struct {
	ctx context.Context

	timeline *tview.TextView
	trend    *tview.TextView
	follow   *tview.TextView
	follower *tview.TextView
	profile  *tview.TextView
	setting  *tview.TextView
	main     *tview.TextView
	footer   *tview.TextView

	horizontalFlex *tview.Flex
	verticalFlex   *tview.Flex

	honeycomb *di.HoneyComb
	app       *tview.Application
	viewModel *viewModel
}

// viewModel is the view model.
type viewModel struct {
	author      *model.Author
	follows     *model.Follows
	timeline    []*model.Post
	currentView *currentView
}

// NewTUI creates a new TUI.
func NewTUI(ctx context.Context, hc *di.HoneyComb) *TUI {
	tui := &TUI{
		ctx:       ctx,
		timeline:  initTimelineTextView(),
		trend:     initTrendTextView(),
		follow:    initFollowTextView(),
		follower:  initFollowerTextView(),
		profile:   initProfileTextView(),
		setting:   initSettingTextView(),
		main:      initMainTextView(),
		footer:    initFooterTextView(),
		honeycomb: hc,
		app:       tview.NewApplication(),
	}

	tui.horizontalFlex = tview.NewFlex().
		AddItem(tui.timeline, 0, 1, false).
		AddItem(tui.trend, 0, 1, false).
		AddItem(tui.follow, 0, 1, false).
		AddItem(tui.follower, 0, 1, false).
		AddItem(tui.profile, 0, 1, false).
		AddItem(tui.setting, 0, 1, false)

	tui.verticalFlex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tui.horizontalFlex, 3, 1, false).
		AddItem(tui.main, 0, 4, false).
		AddItem(tui.footer, 1, 1, false)

	tui.app.SetInputCapture(tui.keyBindings)
	return tui
}

// Run starts the TUI.
func (t *TUI) Run() error {
	if err := t.initializeViewModel(); err != nil {
		return err
	}
	if err := t.writePosts(); err != nil {
		return err
	}

	return t.app.SetRoot(t.verticalFlex, true).Run()
}

func (t *TUI) keyBindings(event *tcell.EventKey) *tcell.EventKey {
	row, column := t.main.GetScrollOffset()
	switch event.Key() {
	case tcell.KeyEsc, tcell.KeyCtrlC:
		t.app.Stop()
	case tcell.KeyTab:
		t.viewModel.currentView.next()
		if err := t.updateViews(); err != nil {
			// TODO:  error handling
			return event
		}
	case tcell.KeyBacktab:
		t.viewModel.currentView.prev()
		if err := t.updateViews(); err != nil {
			// TODO:  error handling
			return event
		}
	case tcell.KeyUp:
		t.main.ScrollTo(row-1, column)
	case tcell.KeyDown:
		t.main.ScrollTo(row+1, column)
	case tcell.KeyPgUp:
		t.main.ScrollTo(row-10, column)
	case tcell.KeyPgDn:
		t.main.ScrollTo(row+10, column)
	}

	switch event.Rune() {
	case 'q':
		t.app.Stop()
	case 'j':
		t.main.ScrollTo(row+1, column)
	case 'k':
		t.main.ScrollTo(row-1, column)
	}
	return event
}

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
		if err := t.writePosts(); err != nil {
			return err
		}
	case currentViewTrend:
		t.main.SetText("Trend")
	case currentViewFollow:
		t.main.SetText("Follow")
	case currentViewFollower:
		t.main.SetText("Follower")
	case currentViewProfile:
		t.main.SetText("Profile")
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
		t.footer.SetText("  Quit:<ESC>, 'q' | Next:<TAB> | Prev:<SHIFT-TAB>")
	case currentViewTrend:
		t.footer.SetText("  Quit:<ESC>, 'q' | Next:<TAB> | Prev:<SHIFT-TAB>")
	case currentViewFollow:
		t.footer.SetText("  Quit:<ESC>, 'q' | Next:<TAB> | Prev:<SHIFT-TAB>")
	case currentViewFollower:
		t.footer.SetText("  Quit:<ESC>, 'q' | Next:<TAB> | Prev:<SHIFT-TAB>")
	case currentViewProfile:
		t.footer.SetText("  Quit:<ESC>, 'q' | Next:<TAB> | Prev:<SHIFT-TAB>")
	case currentViewSetting:
		t.footer.SetText("  Quit:<ESC>, 'q' | Next:<TAB> | Prev:<SHIFT-TAB>")
	}
}

// writePosts writes posts to the main text view.
func (t *TUI) writePosts() error {
	for _, post := range t.viewModel.timeline {
		displayName := post.Author.DisplayName
		content := post.Content
		if _, err := t.main.Write([]byte(fmt.Sprintf("[yellow]%s[white]\n%s\n\n", displayName, content))); err != nil {
			return err
		}
	}
	return nil
}
