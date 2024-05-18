// Package tui provides a text-based user interface for the honeycomb command.
package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nao1215/honeycomb/config"
)

// Run starts the TUI.
func Run() error {
	_, err := config.ReadPrivateKey()
	if err != nil {
		p := tea.NewProgram(newPrivateKeyInputModel())
		if _, err := p.Run(); err != nil {
			return err
		}
	}

	fmt.Println("work in progress...🐝")
	return nil
}
