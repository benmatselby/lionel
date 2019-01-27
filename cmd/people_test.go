package cmd_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"github.com/benmatselby/lionel/cmd"
	"github.com/benmatselby/lionel/mock_trello"
	"github.com/benmatselby/lionel/trello"

	"github.com/golang/mock/gomock"
)

func TestNewListCardsPeopleCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mock_trello.NewMockAPI(ctrl)

	cmd := cmd.NewListCardsPeopleCommand(client)

	use := "people"
	short := "List all the cards for a given board sorted by the people the card is assigned to"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

func TestDisplayCardsForPeople(t *testing.T) {
	tt := []struct {
		name     string
		output   string
		boardErr error
		cardsErr error
		listsErr error
	}{
		{name: "can return a list of cards", output: `Joe Cocker
==========
Points: 32
Cards: 3
- (12) Fire it up
- (20) Find a place where we belong
- Find out if I sang out of tune, what my friends would do

`, boardErr: nil},
		{name: "returns error if board cannot be found", output: "", boardErr: errors.New("something")},
		{name: "returns error if there is an issue returning the cards", output: "", cardsErr: errors.New("something")},
		{name: "returns error if there is an issue returning the lists", output: "", listsErr: errors.New("something")},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := mock_trello.NewMockAPI(ctrl)

			boardJson := `{
"id": "123",
"name": "Golang the musical",
"desc": "Live the dream",
"closed": false,
"memberships": [{
	"idMember": "776655"
}]
}
`
			board := trello.Board{}
			json.Unmarshal([]byte(boardJson), &board)

			cards := []trello.Card{
				{
					Name:          "(12) Fire it up",
					ListID:        "1",
					MembershipIDs: []string{"776655"},
				},
				{
					Name:          "(20) Find a place where we belong",
					ListID:        "1",
					MembershipIDs: []string{"776655"},
				},
				{
					Name:          "Find out if I sang out of tune, what my friends would do",
					ListID:        "1",
					MembershipIDs: []string{"776655"},
				},
			}

			member := trello.Member{
				FullName: "Joe Cocker",
			}

			client.
				EXPECT().
				GetBoard("Golang the musical").
				Return(&board, tc.boardErr).
				AnyTimes()

			client.
				EXPECT().
				GetCards(board).
				Return(cards, tc.cardsErr).
				AnyTimes()

			client.
				EXPECT().
				GetMember("776655").
				Return(&member, tc.listsErr).
				AnyTimes()

			var b bytes.Buffer
			writer := bufio.NewWriter(&b)

			opts := cmd.ListCardPeopleOptions{
				Args: []string{
					"Golang the musical",
				},
			}

			cmd.DisplayCardsForPeople(client, opts, writer)
			writer.Flush()

			if b.String() != tc.output {
				t.Fatalf("expected '%s'; got '%s'", tc.output, b.String())
			}
		})
	}
}
