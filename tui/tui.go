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
	author   *model.Author
	follows  *model.Follows
	timeline []*model.Post
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

	tui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc || event.Key() == tcell.KeyCtrlC || event.Rune() == 'q' {
			tui.app.Stop()
		}
		row, column := tui.main.GetScrollOffset()
		switch event.Key() {
		case tcell.KeyUp:
			tui.main.ScrollTo(row-1, column)
		case tcell.KeyDown:
			tui.main.ScrollTo(row+1, column)
		case tcell.KeyPgUp:
			tui.main.ScrollTo(row-10, column)
		case tcell.KeyPgDn:
			tui.main.ScrollTo(row+10, column)
		}
		return event
	})
	return tui
}

// Run starts the TUI.
func (t *TUI) Run() error {
	if err := t.reloadViewModel(); err != nil {
		return err
	}
	if err := t.writePosts(); err != nil {
		return err
	}

	return t.app.SetRoot(t.verticalFlex, true).Run()
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
