package board

import (
	"fmt"
	"io"
	"os"

	"github.com/benmatselby/lionel/trello"
	"github.com/spf13/cobra"
)

// ListOptions provides the flags for the `list` command
type ListOptions struct {
	ShowClosed bool
	refs       []string
}

// NewListCommand creates a new `board list` command
func NewListCommand(client trello.API) *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all the boards",
		Run: func(cmd *cobra.Command, args []string) {
			opts.refs = args
			err := DisplayBoards(client, opts, os.Stdout)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&opts.ShowClosed, "show-closed", false, "Display closed boards?")

	return cmd
}

// DisplayBoards will render the boards you have access to
func DisplayBoards(client trello.API, opts ListOptions, w io.Writer) error {
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
