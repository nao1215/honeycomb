// Package cmd define subcommands provided by the honeycomb command
package cmd

import (
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "honeycomb",
		Short: `🐝 nostr client CLI application for cross-platform`,
		Long:  `🐝 nostr client CLI application for cross-platform.`,
	}
	cmd.CompletionOptions.DisableDefaultCmd = true
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	cmd.AddCommand(newVersionCmd())
	return cmd
}

// Execute run honeycomb process.
func Execute() error {
	rootCmd := newRootCmd()
	return rootCmd.Execute()
}
