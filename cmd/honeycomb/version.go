package cmd

import (
	"github.com/nao1215/honeycomb/version"
	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:               "version",
		Short:             "Show honeycomb application version",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Printf("honeycomb version %s (MIT LICENSE)\n", version.GetVersion())
		},
	}
}
