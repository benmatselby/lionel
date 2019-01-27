package cmd

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/benmatselby/lionel/trello"
	"github.com/spf13/cobra"
)

// ListCardPeopleOptions provides the flags for the `people` command
type ListCardPeopleOptions struct {
	Args []string
}

// NewListCardsPeopleCommand creates a new `people` command that will show all the cards for a given board
// displayed by assigned to
func NewListCardsPeopleCommand(client trello.API) *cobra.Command {
	var opts ListCardPeopleOptions

	cmd := &cobra.Command{
		Use:   "people",
		Short: "List all the cards for a given board sorted by the people the card is assigned to",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Args = args
			return DisplayCardsForPeople(client, opts, os.Stdout)
		},
	}

	return cmd
}

// Person represents a person on the body, and what cards they have
type Person struct {
	ID         string
	FullName   string
	Cards      []trello.Card
	CardCount  int
	PointCount int
}

// DisplayCardsForPeople will render the boards you have access to
func DisplayCardsForPeople(client trello.API, opts ListCardPeopleOptions, w io.Writer) error {
	board, err := client.GetBoard(opts.Args[0])
	if err != nil {
		return err
	}

	cards, err := client.GetCards(*board)
	if err != nil {
		return err
	}

	var people []Person
	for _, member := range board.Memberships {
		person := Person{
			ID:         member.MemberID,
			CardCount:  0,
			PointCount: 0,
		}

		for _, card := range cards {
			for _, memberID := range card.MembershipIDs {
				if memberID == member.MemberID {
					person.Cards = append(person.Cards, card)
					person.CardCount++

					re := regexp.MustCompile(`^\(.+?\)`)
					pointString := re.FindString(card.Name)
					points, err := strconv.Atoi(strings.Trim(pointString, "()"))
					if err != nil {
						points = 0
					}
					person.PointCount += points
				}
			}
		}

		// Discard anyone without a card
		if person.CardCount > 0 {
			user, err := client.GetMember(member.MemberID)
			if err != nil {
				return err
			}
			person.FullName = user.FullName
			people = append(people, person)
		}
	}

	for _, person := range people {
		fmt.Fprintf(w, "%s\n%s\n", person.FullName, strings.Repeat("=", len(person.FullName)))
		fmt.Fprintf(w, "Points: %v\n", person.PointCount)
		fmt.Fprintf(w, "Cards: %v\n", person.CardCount)
		for _, card := range person.Cards {
			fmt.Fprintf(w, "- %s\n", card.Name)
		}
		fmt.Fprintln(w)
	}

	return nil
}
