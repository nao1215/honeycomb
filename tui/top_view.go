// Package tui provides a text-based user interface for the honeycomb command.
package tui

import (
	"context"

	"github.com/nao1215/honeycomb/app/di"
)

// Run starts the TUI.
func Run() error {
	ctx := context.Background()
	honeyComb, err := di.NewHoneyComb(ctx)
	if err != nil {
		return err
	}
	tui := NewTUI(ctx, honeyComb)
	return tui.Run()
}
