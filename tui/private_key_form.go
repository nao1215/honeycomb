package tui

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/nao1215/honeycomb/app/model"
	"github.com/rivo/tview"
	"golang.design/x/clipboard"
)

// showNSecretKeyForm shows the NSecretKey form and writes the input NSecretKey to a file.
func showNSecretKeyForm() (model.NSecretKey, error) {
	app := tview.NewApplication()
	form := tview.NewForm()
	form.SetBorder(true).SetTitle("🐝  NSecretKey  🐝").SetTitleAlign(tview.AlignCenter)
	form.AddInputField("NSecretKey(nsec**): ", "", 60, nil, nil)

	nsecretKey := model.NSecretKey("")
	form.AddButton("Submit", func() {
		inputField, ok := form.GetFormItem(0).(*tview.InputField)
		if !ok {
			showError(app, "Failed to read NSecretKey input field.")
			return
		}

		nsecretKey = model.NSecretKey(inputField.GetText())
		if err := nsecretKey.Validate(); err != nil {
			showError(app, err.Error())
			return
		}
		if err := model.WriteNSecretKey(nsecretKey); err != nil {
			showError(app, err.Error())
			return
		}
		app.Stop()
	})

	form.AddButton("Cancel", func() {
		app.Stop()
	})

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlC, tcell.KeyEsc:
			app.Stop()
			return nil
		default:
			return event
		}
	})

	form.SetMouseCapture(func(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {
		if action == tview.MouseRightClick {
			err := clipboard.Init()
			if err != nil {
				showError(app, err.Error())
				return action, event
			}

			clipText := strings.TrimSpace(string(clipboard.Read(clipboard.FmtText)))
			inputField, ok := form.GetFormItem(0).(*tview.InputField)
			if !ok {
				showError(app, "Failed to read NSecretKey input field.")
				return action, event
			}
			inputField.SetText(clipText)
			return tview.MouseConsumed, nil
		}
		return action, event
	})

	app.EnableMouse(true)

	if err := app.SetRoot(form, true).Run(); err != nil {
		return "", err
	}
	return nsecretKey, nil
}

// showError displays an error message.
func showError(app *tview.Application, message string) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(_ int, _ string) {
			app.Stop()
		})
	app.SetRoot(modal, false).SetFocus(modal)
}
