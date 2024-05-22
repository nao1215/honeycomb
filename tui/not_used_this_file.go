package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
	"github.com/nao1215/honeycomb/app/model"
)

// status is the status of the honeycomb operation.
type status uint

const (
	// statusNone is the status when the honeycomb operation is not executed.
	statusNone status = iota
	// statusPrivateKeyInput is the status when the private key input view is displayed.
	statusPrivateKeyInput
	// statusPrivateKeyValidateErr is the status when the private key validation error occurs.
	statusPrivateKeyValidateErr
	// statusPrivateKeySaveErr is the status when the private key save error occurs.
	statusPrivateKeySaveErr
)

// subtle returns a string with a subtle color.
func subtle(message string) string {
	return termenv.String(message).Foreground(termenv.ColorProfile().Color("241")).String()
}

// red returns a string with a red color.
func red(message string) string {
	return termenv.String(message).Foreground(termenv.ColorProfile().Color("196")).String()
}

// green returns a string with a green color.
func green(message string) string {
	return termenv.String(message).Foreground(termenv.ColorProfile().Color("46")).String()
}

// privateKeyInputModel is the model for the private key input view.
type privateKeyInputModel struct {
	textInput textinput.Model // text input model for getting the private key
	status    status          // status of the honeycomb operation
	err       error           // error occurred during the input
}

// newPrivateKeyInputModel creates a new privateKeyInputModel.
func newPrivateKeyInputModel() *privateKeyInputModel {
	ti := textinput.New()
	ti.Placeholder = "nsec-..."
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 64

	return &privateKeyInputModel{
		textInput: ti,
		status:    statusNone,
		err:       nil,
	}
}

// Init initializes the private key input view.
func (m *privateKeyInputModel) Init() tea.Cmd {
	m.status = statusPrivateKeyInput
	return textinput.Blink
}

// Update updates the private key input view.
func (m *privateKeyInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg: //nolint:exhaustive
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.status == statusPrivateKeySaveErr {
				return m, tea.Quit
			}

			if m.textInput.Value() == "" {
				m.status = statusPrivateKeyValidateErr
				m.err = fmt.Errorf("private key is empty")
				return m, nil
			}

			pk := model.NSecretKey(m.textInput.Value())
			if err := pk.Validate(); err != nil {
				m.status = statusPrivateKeyValidateErr
				m.err = err
				return m, nil
			}

			if err := model.WriteNSecretKey(pk); err != nil {
				m.status = statusPrivateKeySaveErr
				m.err = err
			}
			// TODO: move next view.
			return m, tea.Quit
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// View returns the private key input view.
func (m *privateKeyInputModel) View() string {
	switch m.status { //nolint:exhaustive
	case statusPrivateKeyValidateErr:
		text := fmt.Sprintf(
			"%s\n%s%s\n\n%s\n%s\n%s\n",
			"🐝 Please input a private key that starts with 'nsec'.",
			"🐝 The private key will be saved to ", model.NSecretKeyFilePath(),
			m.textInput.View(),
			red(fmt.Sprintf("  Warning: %s", m.err.Error())),
			subtle("ESC or <Ctrl-C>:quit  Enter:submit"),
		)
		return text
	case statusPrivateKeySaveErr:
		return red(fmt.Sprintf("can not save private key: %s", m.err.Error()))
	default:
		return fmt.Sprintf(
			"%s\n%s%s\n\n%s\n\n%s\n",
			"🐝 Please input a private key that starts with 'nsec'.",
			"🐝 The private key will be saved to ", model.NSecretKeyFilePath(),
			m.textInput.View(),
			subtle("ESC or <Ctrl-C>:quit  Enter:submit"),
		)
	}
}
