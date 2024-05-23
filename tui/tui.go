package tui

import (
	"context"

	"github.com/gdamore/tcell/v2"
	"github.com/nao1215/honeycomb/app/di"
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

	postForm        *tview.Form
	postFormVisible *postFormVisible

	honeycomb *di.HoneyComb
	app       *tview.Application
	viewModel *viewModel
}

// keyBindings handles key bindings.
func (t *TUI) keyBindings(event *tcell.EventKey) *tcell.EventKey {
	if t.postFormVisible.isVisible() {
		switch event.Key() {
		case tcell.KeyEsc:
			t.postFormVisible.invisible()
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
	case 'p':
		t.postFormVisible.visible()
		t.app.SetRoot(t.postForm, true).EnableMouse(true)
	}
	return event
}

// mouseHandler handles mouse events in the main text view.
func (t *TUI) mouseHandler(event *tcell.EventMouse, action tview.MouseAction) (*tcell.EventMouse, tview.MouseAction) {
	switch *t.viewModel.currentView {
	case currentViewTimeline:
		return t.timelineMouseHandler(event, action)
	default:
		return event, action
	}
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
				// clickedPost := postRange.post
				// TODO: Handle the clicked post as needed
				return nil, action // consume the event
			}
		}
	}
	return event, action
}
