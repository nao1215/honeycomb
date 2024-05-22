package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gdamore/tcell/v2"
	"github.com/nao1215/honeycomb/app/model"
	"github.com/nao1215/honeycomb/app/usecase"
	"github.com/rivo/tview"
	"github.com/shogo82148/pointer"
)

// initializeViewModel reloads the view model.
func (t *TUI) initializeViewModel() error {
	nsecKey, err := model.ReadNSecretKey()
	if err != nil {
		// TODO: rewrite to use tview
		p := tea.NewProgram(newPrivateKeyInputModel())
		if _, err := p.Run(); err != nil {
			return err
		}
		return nil
	}

	author, err := t.honeycomb.GetAuthor(t.ctx, &usecase.AuthorGetterInput{
		NSecretKey: nsecKey,
	})
	if err != nil {
		return err
	}
	defer author.Author.Close() //nolint:errcheck

	follows, err := t.honeycomb.ListFollow(t.ctx, &usecase.FollowListerInput{
		PublicKey:      author.Author.PublicKey,
		ConnectedRelay: author.Author.ConnectedRelay,
	})
	if err != nil {
		return err
	}

	timeline, err := t.honeycomb.ListTimeline(t.ctx, &usecase.TimelineListerInput{
		Follows:        follows.Follows,
		Limit:          100,
		ConnectedRelay: author.Author.ConnectedRelay,
	})
	if err != nil {
		return err
	}

	t.viewModel = &viewModel{
		author:      author.Author,
		follows:     &follows.Follows,
		timeline:    timeline.Posts,
		currentView: pointer.Ptr(currentViewTimeline),
	}
	return nil
}

// initTimelineTextView initializes the timeline text view.
func initTimelineTextView() *tview.TextView {
	timeline := tview.NewTextView()
	timeline.SetBorder(true).SetBorderColor(tcell.ColorWhite)
	timeline.SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite).SetText("🐝  TL  🐝")
	return timeline
}

// initTrendTextView initializes the trend text view.
func initTrendTextView() *tview.TextView {
	trend := tview.NewTextView()
	trend.SetBorder(true).SetBorderColor(tcell.ColorWhite)
	trend.SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite).SetText("Trend")
	return trend
}

// initFollowTextView initializes the follow text view.
func initFollowTextView() *tview.TextView {
	follow := tview.NewTextView()
	follow.SetBorder(true).SetBorderColor(tcell.ColorWhite)
	follow.SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite).SetText("Follow")
	return follow
}

// initFollowerTextView initializes the follower text view.
func initFollowerTextView() *tview.TextView {
	follower := tview.NewTextView()
	follower.SetBorder(true).SetBorderColor(tcell.ColorWhite)
	follower.SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite).SetText("Follower")
	return follower
}

// initProfileTextView initializes the profile text view.
func initProfileTextView() *tview.TextView {
	profile := tview.NewTextView()
	profile.SetBorder(true).SetBorderColor(tcell.ColorWhite)
	profile.SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite).SetText("Profile")
	return profile
}

// initSettingTextView initializes the setting text view.
func initSettingTextView() *tview.TextView {
	setting := tview.NewTextView()
	setting.SetBorder(true).SetBorderColor(tcell.ColorWhite)
	setting.SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite).SetText("Setting")
	return setting
}

// initMainTextView initializes the main text view.
func initMainTextView() *tview.TextView {
	main := tview.NewTextView()
	main.SetBorder(true).SetBorderColor(tcell.ColorWhite)
	main.SetTextAlign(tview.AlignLeft).SetTextColor(tcell.ColorWhite)
	main.SetDynamicColors(true).SetScrollable(true)
	return main
}

// initFooterTextView initializes the footer text view.
func initFooterTextView() *tview.TextView {
	footer := tview.NewTextView()
	footer.SetTextAlign(tview.AlignLeft).SetTextColor(tcell.ColorLightGoldenrodYellow).SetText("  Quit:<ESC>, 'q'")
	return footer
}
