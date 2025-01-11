package cmd

import (
	"github.com/spf13/cobra"
)

func addMemberCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "add",
		Short:   "for adding members",
		Long:    "Add member(s) to the list of members",
		Args:    cobra.MinimumNArgs(1),
		Example: `add Foo Bar Baz`,
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, member := range args {
				cmd.Printf("Adding member: %s\n", member)
			}

			return nil
		},
	}
}
