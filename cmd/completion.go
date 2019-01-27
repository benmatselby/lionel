package cmd

import (
	"github.com/spf13/cobra"
)

// NewCompletionCommand creates a new `completion` command
func NewCompletionCommand(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generates the bash completion script: lionel.sh",
		Run: func(cmd *cobra.Command, args []string) {
			root.GenBashCompletionFile("lionel.sh")
		},
	}
	return cmd
}
