package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "v0.0.1"
var rootCmd = &cobra.Command{
	Use:     "fair-chance-go",
	Version: version,
	Short:   "A CLI app to fairly select a name from a list of names",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	rootCmd.AddCommand(addMemberCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
