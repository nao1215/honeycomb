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

	// TODO: Delete under code.
	ctx := context.Background()
	honeyComb, err := di.NewHoneyComb(ctx)
	if err != nil {
		return err
	}

	authorOutput, err := honeyComb.GetAuthor(ctx, &usecase.AuthorGetterInput{
		NSecretKey: nsecKey,
	})
	if err != nil {
		return err
	}
	defer authorOutput.Author.Close() //nolint:errcheck

	follow, err := honeyComb.ListFollow(ctx, &usecase.FollowListerInput{
		PublicKey:      authorOutput.Author.PublicKey,
		ConnectedRelay: authorOutput.Author.ConnectedRelay,
	})
	if err != nil {
		return err
	}

	fmt.Println("[WebSite]")
	fmt.Printf("%s\n", authorOutput.Author.Profile.Website)
	fmt.Println("[Nip05]")
	fmt.Printf("%s\n", authorOutput.Author.Profile.Nip05)
	fmt.Println("[Picture]")
	fmt.Printf("%s\n", authorOutput.Author.Profile.Picture)
	fmt.Println("[Lud16]")
	fmt.Printf("%s\n", authorOutput.Author.Profile.Lud16)
	fmt.Println("[DisplayName]")
	fmt.Printf("%s\n", authorOutput.Author.Profile.DisplayName)
	fmt.Println("[About]")
	fmt.Printf("%s\n", authorOutput.Author.Profile.About)
	fmt.Println("[Name]")
	fmt.Printf("%s\n", authorOutput.Author.Profile.Name)
	fmt.Println("[Bot]")
	fmt.Printf("%t\n", authorOutput.Author.Profile.Bot)
	fmt.Println("[NpublicKey]")
	fmt.Printf("%s\n\n", authorOutput.Author.NPublicKey)

	for i, v := range follow.Follows {
		fmt.Printf("[Follow:%d]\n", i+1)
		fmt.Printf("DisplayName:%s\n", v.Profile.DisplayName)
		fmt.Printf("Public Key:%s\n", v.PublicKey)
	}
	fmt.Println()
	fmt.Println("work in progress...🐝")
	return nil
}
