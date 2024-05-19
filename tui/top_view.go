// Package tui provides a text-based user interface for the honeycomb command.
package tui

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nao1215/honeycomb/app/di"
	"github.com/nao1215/honeycomb/app/model"
	"github.com/nao1215/honeycomb/app/usecase"
)

// Run starts the TUI.
func Run() error {
	nsecKey, err := model.ReadNSecretKey()
	if err != nil {
		p := tea.NewProgram(newPrivateKeyInputModel())
		if _, err := p.Run(); err != nil {
			return err
		}
		return nil
	}

	ctx := context.Background()
	honeyComb, err := di.NewHoneyComb(ctx)
	if err != nil {
		return err
	}

	output, err := honeyComb.ProfileGetter.GetProfile(ctx, &usecase.ProfileGetterInput{
		NsecretKey: nsecKey,
	})
	if err != nil {
		return err
	}

	fmt.Println("[WebSite]")
	fmt.Printf("%s\n", output.Profile.Website)
	fmt.Println("[Nip05]")
	fmt.Printf("%s\n", output.Profile.Nip05)
	fmt.Println("[Picture]")
	fmt.Printf("%s\n", output.Profile.Picture)
	fmt.Println("[Lud16]")
	fmt.Printf("%s\n", output.Profile.Lud16)
	fmt.Println("[DisplayName]")
	fmt.Printf("%s\n", output.Profile.DisplayName)
	fmt.Println("[About]")
	fmt.Printf("%s\n", output.Profile.About)
	fmt.Println("[Name]")
	fmt.Printf("%s\n", output.Profile.Name)
	fmt.Println("[Bot]")
	fmt.Printf("%t\n", output.Profile.Bot)
	fmt.Println("[NpublicKey]")
	fmt.Printf("%s\n", output.NpublicKey)

	fmt.Println("work in progress...🐝")
	return nil
}
