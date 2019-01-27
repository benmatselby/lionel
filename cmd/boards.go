package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/benmatselby/lionel/trello"
	"github.com/spf13/cobra"
)

// ListBoardOptions provides the flags for the `list` command
type ListBoardOptions struct {
	ShowClosed bool
	refs       []string
}

// NewListBoardsCommand creates a new `board list` command
func NewListBoardsCommand(client trello.API) *cobra.Command {
	var opts ListBoardOptions

	cmd := &cobra.Command{
		Use:   "boards",
		Short: "List all the boards",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.refs = args
			return DisplayBoards(client, opts, os.Stdout)
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&opts.ShowClosed, "show-closed", false, "Display closed boards?")

	return cmd
}

// DisplayBoards will render the boards you have access to
func DisplayBoards(client trello.API, opts ListBoardOptions, w io.Writer) error {
	boards, err := client.GetBoards()
	if err != nil {
		return err
	}

	for _, board := range boards {
		if board.Closed != true || (board.Closed && opts.ShowClosed) {
			fmt.Fprintln(w, board.Name)
		}
	}

	return nil
}
