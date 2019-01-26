package board

import (
	"github.com/benmatselby/lionel/trello"
	"github.com/spf13/cobra"
)

// NewBoardCommand creates a new `board` command
func NewBoardCommand(client trello.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "board",
		Short: "Commands that relate to Trello boards",
	}

	cmd.AddCommand(
		NewListCommand(client),
	)

	return cmd
}
