package cmd

import (
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "cli",
		Short: "CLI for running tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
