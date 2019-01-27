package cmd

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/benmatselby/lionel/trello"
	"github.com/spf13/cobra"
)

// BurndownOptions provides the flags for the `cards` command
type BurndownOptions struct {
	Args []string
}

// NewBurndownCommand creates a new `burndown` command
func NewBurndownCommand(client trello.API) *cobra.Command {
	var opts BurndownOptions
	cmd := &cobra.Command{
		Use:   "burndown",
		Short: "Provide a burndown table using the 'scrum for trello' plugin data",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Args = args
			return DisplayBurndown(client, opts, os.Stdout)
		},
	}

	return cmd
}

// DisplayBurndown will render burndown data using the story points associated with the card
// when you use the scrum for trello plugin
func DisplayBurndown(client trello.API, opts BurndownOptions, w io.Writer) error {
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

	tw := tabwriter.NewWriter(w, 0, 1, 1, ' ', 0)
	fmt.Fprintf(tw, "%s\t%s\t%s\n", "List", "Cards", "Story Points")
	fmt.Fprintf(tw, "%s\t%s\t%s\n", "----", "-----", "------------")

	totalCard := 0
	totalPoint := 0
	for _, list := range lists {
		cardCount := 0
		pointCount := 0
		for _, card := range cards {
			if card.ListID == list.ID {
				cardCount++

				re := regexp.MustCompile(`^\(.+?\)`)
				pointString := re.FindString(card.Name)
				points, err := strconv.Atoi(strings.Trim(pointString, "()"))
				if err != nil {
					points = 0
				}
				pointCount += points
			}
		}
		totalCard += cardCount
		totalPoint += pointCount
		fmt.Fprintf(tw, "%s\t%v\t%v\n", list.Name, totalCard, totalPoint)
	}

	fmt.Fprintf(tw, "%s\t%s\t%s\n", "-----", "-----", "------------")
	fmt.Fprintf(tw, "%s\t%d\t%d\n", "Total", totalCard, totalPoint)
	fmt.Fprintf(tw, "%s\t%s\t%s\n", "-----", "-----", "------------")

	tw.Flush()

	return nil
}
