package cmd_test

import (
	"bufio"
	"bytes"
	"errors"
	"testing"

	"github.com/benmatselby/lionel/cmd"
	"github.com/benmatselby/lionel/mock_trello"
	"github.com/benmatselby/lionel/trello"

	"github.com/golang/mock/gomock"
)

func TestNewBurndownCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mock_trello.NewMockAPI(ctrl)

	cmd := cmd.NewBurndownCommand(client)

	use := "burndown"
	short := "Provide a burndown table using the 'scrum for trello' plugin data"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

func TestDisplayBurndown(t *testing.T) {
	tt := []struct {
		name     string
		output   string
		boardErr error
		cardsErr error
		listsErr error
	}{
		{name: "can return a list of cards", output: `List     Cards Story Points
----     ----- ------------
To do    2     1
Progress 1     250
-----    ----- ------------
Total    3     251
-----    ----- ------------
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

			board := trello.Board{
				ID:   "1",
				Name: "Golang the musical",
			}

			cards := []trello.Card{
				{
					Name:   "(250) Write a musical",
					ListID: "2",
				},
				{
					Name:   "(1) Perform the musical",
					ListID: "1",
				},
				{
					Name:   "Paint a picture, just like how it use to be",
					ListID: "1",
				},
			}

			lists := []trello.List{
				{
					ID:   "1",
					Name: "To do",
				},
				{
					ID:   "2",
					Name: "Progress",
				},
			}

			client.
				EXPECT().
				GetBoard("a board").
				Return(&board, tc.boardErr).
				AnyTimes()

			client.
				EXPECT().
				GetCards(board).
				Return(cards, tc.cardsErr).
				AnyTimes()

			client.
				EXPECT().
				GetLists(board).
				Return(lists, tc.listsErr).
				AnyTimes()

			var b bytes.Buffer
			writer := bufio.NewWriter(&b)

			opts := cmd.BurndownOptions{
				Args: []string{
					"a board",
				},
			}

			cmd.DisplayBurndown(client, opts, writer)
			writer.Flush()

			if b.String() != tc.output {
				t.Fatalf("expected '%s'; got '%s'", tc.output, b.String())
			}
		})
	}
}
