// Package tui provides the text-based user interface.
package tui

import (
	"context"

	"github.com/gdamore/tcell/v2"
	"github.com/nao1215/honeycomb/app/di"
	"github.com/nao1215/honeycomb/app/usecase"
	"github.com/rivo/tview"
	"github.com/shogo82148/pointer"
)

// TUI is the text-based user interface.
type TUI struct {
	ctx context.Context

	timeline *tview.TextView
	follow   *tview.TextView
	profile  *tview.TextView
	main     *tview.TextView
	footer   *tview.TextView

	horizontalFlex *tview.Flex
	verticalFlex   *tview.Flex

	postForm         *tview.Form
	postFormVisible  *visible
	postModalVisible *visible

	honeycomb *di.HoneyComb
	app       *tview.Application
	viewModel *viewModel
}

// keyBindings handles key bindings.
func (t *TUI) keyBindings(event *tcell.EventKey) *tcell.EventKey {
	if t.postFormVisible.isVisible() || t.postModalVisible.isVisible() {
		switch event.Key() {
		case tcell.KeyEsc:
			t.postFormVisible.invisible()
			t.postModalVisible.invisible()
			t.app.SetRoot(t.verticalFlex, true)
			return nil
		default:
			return event
		}
	}

	row, column := t.main.GetScrollOffset()
	switch event.Key() {
	case tcell.KeyEsc, tcell.KeyCtrlC:
		t.app.Stop()
	case tcell.KeyTab:
		t.viewModel.currentView.next()
		if err := t.updateViews(); err != nil {
			showError(t.app, err.Error())
			return event
		}
	case tcell.KeyBacktab:
		t.viewModel.currentView.prev()
		if err := t.updateViews(); err != nil {
			showError(t.app, err.Error())
			return event
		}
	case tcell.KeyUp:
		t.main.ScrollTo(row-1, column)
	case tcell.KeyDown:
		t.main.ScrollTo(row+1, column)
		if err := t.appendOldTimelineAndRewriteIfNeeded(); err != nil {
			showError(t.app, err.Error())
			return event
		}
	case tcell.KeyPgUp:
		t.main.ScrollTo(row-10, column)
	case tcell.KeyPgDn:
		t.main.ScrollTo(row+10, column)
		if err := t.appendOldTimelineAndRewriteIfNeeded(); err != nil {
			showError(t.app, err.Error())
			return event
		}
	case tcell.KeyEnter:
		if *t.viewModel.currentView == currentViewTimeline {
			t.timelineMouseHandler(tcell.NewEventMouse(
				0, // dummy
				headerOffsetLineCount,
				tcell.Button1, // dummy
				tcell.ModNone, // dummy
			), tview.MouseLeftClick)
		}
	}

	switch event.Rune() {
	case 'q':
		t.app.Stop()
	case 'j':
		t.main.ScrollTo(row+1, column)
		if err := t.appendOldTimelineAndRewriteIfNeeded(); err != nil {
			showError(t.app, err.Error())
			return event
		}
	case 'k':
		t.main.ScrollTo(row-1, column)
	case 'p':
		t.postFormVisible.visible()
		t.app.SetRoot(t.postForm, true).EnableMouse(true)
	case 'R':
		if err := t.reload(); err != nil {
			showError(t.app, err.Error())
			return event
		}
	}
	return event
}

// reload reloads the viewmodel data.
func (t *TUI) reload() error {
	switch *t.viewModel.currentView {
	case currentViewTimeline:
		timeline, err := t.honeycomb.ListTimeline(t.ctx, &usecase.TimelineListerInput{
			Follows:        *t.viewModel.follows,
			Limit:          15,
			ConnectedRelay: t.viewModel.author.ConnectedRelay,
		})
		if err != nil {
			return err
		}
		t.viewModel.timeline = timeline.Posts

		if err := t.updateViews(); err != nil {
			return err
		}
	case currentViewFollow:
		follows, err := t.honeycomb.ListFollow(t.ctx, &usecase.FollowListerInput{
			PublicKey:      t.viewModel.author.PublicKey,
			ConnectedRelay: t.viewModel.author.ConnectedRelay,
		})
		if err != nil {
			return err
		}
		t.viewModel.follows = pointer.Ptr(follows.Follows)

		if err := t.updateViews(); err != nil {
			return err
		}
	}
	return nil
}

// mouseHandler handles mouse events in the main text view.
func (t *TUI) mouseHandler(event *tcell.EventMouse, action tview.MouseAction) (*tcell.EventMouse, tview.MouseAction) {
	// Previously, you had to right-click a post to react to it. With this specification,
	// it wasn't possible to click on any links within the post.
	// Therefore, this specification has been removed.
	return event, action
}

// timelineMouseHandler handles mouse events in the timeline view.
func (t *TUI) timelineMouseHandler(event *tcell.EventMouse, action tview.MouseAction) (*tcell.EventMouse, tview.MouseAction) {
	if action == tview.MouseLeftClick {
		_, y := event.Position()
		lineOffset, _ := t.main.GetScrollOffset()
		clickedLine := y + lineOffset

		// Find the post that was clicked on
		for _, postRange := range t.viewModel.postRanges {
			if clickedLine >= postRange.startLine && clickedLine < postRange.startLine+postRange.lineCount {
				t.showPostModal(postRange.post)
				return nil, action
			}
		}
	}
	return event, action
}
