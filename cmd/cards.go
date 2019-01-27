package cmd

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/benmatselby/lionel/trello"
	"github.com/spf13/cobra"
)

// ListCardOptions provides the flags for the `list` command
type ListCardOptions struct {
	StripScrumTags bool
	Args           []string
}

// NewListCardsCommand creates a new `board list` command
func NewListCardsCommand(client trello.API) *cobra.Command {
	var opts ListCardOptions

	cmd := &cobra.Command{
		Use:   "cards",
		Short: "List all the cards for a given board",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Args = args
			return DisplayCards(client, opts, os.Stdout)
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&opts.StripScrumTags, "strip-scrum", false, `Remove the scrum for trello plugin brackets.
These are normally () prefixes to the card name, with a number inside`)

	return cmd
}

// DisplayCards will render the boards you have access to
func DisplayCards(client trello.API, opts ListCardOptions, w io.Writer) error {
	board, err := client.GetBoard(opts.Args[0])
	if err != nil {
		return err
	}

	cards, err := client.GetCards(*board)
	if err != nil {
		return err
	}

	lists, err := client.GetLists(*board)
	if err != nil {
		return err
	}

	for _, list := range lists {
		listCards := ""
		listCount := 0
		for _, card := range cards {
			if card.ListID == list.ID {
				listCount++
				cardName := card.Name

				if opts.StripScrumTags == true {
					re := regexp.MustCompile(`^\(.+?\) `)
					cardName = re.ReplaceAllString(cardName, "")
				}
				listCards += fmt.Sprintf("* %s\n", cardName)
			}
		}

		listTitle := fmt.Sprintf("%s (%v)", list.Name, listCount)
		fmt.Fprintf(w, "%s\n%s\n\n", listTitle, strings.Repeat("=", len(listTitle)))
		fmt.Fprintf(w, "%s\n", listCards)
	}

	return nil
}
