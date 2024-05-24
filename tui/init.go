package tui

import (
	"context"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/nao1215/honeycomb/app/di"
	"github.com/nao1215/honeycomb/app/model"
	"github.com/nao1215/honeycomb/app/usecase"
	"github.com/rivo/tview"
	"github.com/shogo82148/pointer"
	"golang.design/x/clipboard"
)

// Run starts the TUI.
func Run() error {
	ctx := context.Background()
	honeyComb, err := di.NewHoneyComb(ctx)
	if err != nil {
		return err
	}
	tui := newTUI(ctx, honeyComb)
	return tui.run()
}

// newTUI creates a new TUI.
func newTUI(ctx context.Context, hc *di.HoneyComb) *TUI {
	tui := &TUI{
		ctx:              ctx,
		timeline:         initTimelineTextView(),
		trend:            initTrendTextView(),
		follow:           initFollowTextView(),
		follower:         initFollowerTextView(),
		profile:          initProfileTextView(),
		setting:          initSettingTextView(),
		main:             initMainTextView(),
		footer:           initFooterTextView(),
		postFormVisible:  pointer.Ptr(visible(false)),
		postModalVisible: pointer.Ptr(visible(false)),
		honeycomb:        hc,
		app:              tview.NewApplication(),
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

	tui.initPostForm()

	tui.app.SetInputCapture(tui.keyBindings)
	tui.app.SetMouseCapture(tui.mouseHandler)
	tui.app.EnableMouse(true)
	tui.app.EnablePaste(true)

	return tui
}

// run starts the TUI.
func (t *TUI) run() error {
	nsecKey, err := model.ReadNSecretKey()
	if err != nil {
		nsecKey, err = showNSecretKeyForm()
		if err != nil {
			return err
		}
		if nsecKey == "" {
			return nil
		}
	}

	if err := t.initializeViewModel(nsecKey); err != nil {
		return err
	}
	defer t.viewModel.author.Close() //nolint:errcheck

	if err := t.writeTimeline(); err != nil {
		return err
	}
	return t.app.SetRoot(t.verticalFlex, true).Run()
}

// initializeViewModel reloads the view model.
func (t *TUI) initializeViewModel(nsecKey model.NSecretKey) error {
	author, err := t.honeycomb.GetAuthor(t.ctx, &usecase.AuthorGetterInput{
		NSecretKey: nsecKey,
	})
	if err != nil {
		return err
	}

	follows, err := t.honeycomb.ListFollow(t.ctx, &usecase.FollowListerInput{
		PublicKey:      author.Author.PublicKey,
		ConnectedRelay: author.Author.ConnectedRelay,
	})
	if err != nil {
		return err
	}

	timeline, err := t.honeycomb.ListTimeline(t.ctx, &usecase.TimelineListerInput{
		Follows:        follows.Follows,
		Limit:          15,
		ConnectedRelay: author.Author.ConnectedRelay,
	})
	if err != nil {
		return err
	}

	myPosts, err := t.honeycomb.ListTimeline(t.ctx, &usecase.TimelineListerInput{
		Follows: model.Follows{{
			PublicKey: author.Author.PublicKey,
			Profile:   *author.Author.Profile,
		}},
		Limit:          3,
		ConnectedRelay: author.Author.ConnectedRelay,
	})
	if err != nil {
		return err
	}

	t.viewModel = &viewModel{
		author:      author.Author,
		follows:     &follows.Follows,
		timeline:    timeline.Posts,
		myPosts:     myPosts.Posts,
		currentView: pointer.Ptr(currentViewTimeline),
	}
	return nil
}

// initTimelineTextView initializes the timeline text view.
func initTimelineTextView() *tview.TextView {
	timeline := tview.NewTextView()
	timeline.SetBorder(true).SetBorderColor(tcell.ColorWhite).SetBackgroundColor(tcell.ColorDefault)
	timeline.SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite).SetText("🐝  TL  🐝")
	return timeline
}

// initTrendTextView initializes the trend text view.
func initTrendTextView() *tview.TextView {
	trend := tview.NewTextView()
	trend.SetBorder(true).SetBorderColor(tcell.ColorWhite).SetBackgroundColor(tcell.ColorDefault)
	trend.SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite).SetText("Trend")
	return trend
}

// initFollowTextView initializes the follow text view.
func initFollowTextView() *tview.TextView {
	follow := tview.NewTextView()
	follow.SetBorder(true).SetBorderColor(tcell.ColorWhite).SetBackgroundColor(tcell.ColorDefault)
	follow.SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite).SetText("Follow")
	return follow
}

// initFollowerTextView initializes the follower text view.
func initFollowerTextView() *tview.TextView {
	follower := tview.NewTextView()
	follower.SetBorder(true).SetBorderColor(tcell.ColorWhite).SetBackgroundColor(tcell.ColorDefault)
	follower.SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite).SetText("Follower")
	return follower
}

// initProfileTextView initializes the profile text view.
func initProfileTextView() *tview.TextView {
	profile := tview.NewTextView()
	profile.SetBorder(true).SetBorderColor(tcell.ColorWhite).SetBackgroundColor(tcell.ColorDefault)
	profile.SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite).SetText("Profile")
	return profile
}

// initSettingTextView initializes the setting text view.
func initSettingTextView() *tview.TextView {
	setting := tview.NewTextView()
	setting.SetBorder(true).SetBorderColor(tcell.ColorWhite).SetBackgroundColor(tcell.ColorDefault)
	setting.SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite).SetText("Setting")
	return setting
}

// initMainTextView initializes the main text view.
func initMainTextView() *tview.TextView {
	main := tview.NewTextView()
	main.SetBorder(true).SetBorderColor(tcell.ColorWhite).SetBackgroundColor(tcell.ColorDefault)
	main.SetTextAlign(tview.AlignLeft).SetTextColor(tcell.ColorWhite)
	main.SetDynamicColors(true).SetScrollable(true)
	return main
}

// initFooterTextView initializes the footer text view.
func initFooterTextView() *tview.TextView {
	footer := tview.NewTextView()
	footer.SetTextAlign(tview.AlignLeft).
		SetTextColor(tcell.ColorLightGoldenrodYellow).
		SetText("  Quit:<ESC> | Next:<TAB> | Prev:<SHIFT-TAB> | Post:'p' | Reload: 'R' | Reaction: <Enter>")
	footer.SetBackgroundColor(tcell.ColorDefault)
	return footer
}

// initPostForm initializes the post form.
func (t *TUI) initPostForm() {
	t.postForm = tview.NewForm().AddTextArea("", "", 100, 10, 1000, nil)
	t.postForm.AddButton("Post", t.writePost)
	t.postForm.AddButton("Cancel", func() {
		t.postFormVisible.invisible()
		t.app.SetRoot(t.verticalFlex, true)
	})

	t.postForm.SetMouseCapture(func(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {
		if action == tview.MouseRightClick {
			err := clipboard.Init()
			if err != nil {
				showError(t.app, err.Error())
				return action, event
			}

			clipText := strings.TrimSpace(string(clipboard.Read(clipboard.FmtText)))
			textArea, ok := t.postForm.GetFormItem(0).(*tview.TextArea)
			if !ok {
				showError(t.app, "Failed to read post input field.")
				return action, event
			}
			textArea.SetText(clipText, true)
			return tview.MouseConsumed, nil
		}
		return action, event
	})
	t.postForm.SetBorder(true).SetTitle("🐝  New Post  🐝").SetTitleAlign(tview.AlignCenter)
}
