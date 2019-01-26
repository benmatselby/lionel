package board_test

import (
	"bufio"
	"bytes"
	"errors"
	"testing"

	"github.com/benmatselby/lionel/cmd/trello/board"
	"github.com/benmatselby/lionel/mock_trello"
	"github.com/benmatselby/lionel/trello"

	"github.com/golang/mock/gomock"
)

func TestNewListCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mock_trello.NewMockAPI(ctrl)

	cmd := board.NewListCommand(client)

	use := "list"
	short := "List all the boards"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

func TestDisplayBoards(t *testing.T) {
	tt := []struct {
		name   string
		output string
		closed bool
		err    error
	}{
		{name: "can return a list of all boards", output: "Magical board\nAdventure board\n", closed: true, err: nil},
		{name: "can return a list of non closed boards", output: "Magical board\n", closed: false, err: nil},
		{name: "returns error if we cannot get list of boards", output: "", err: errors.New("something")},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := mock_trello.NewMockAPI(ctrl)

			boards := trello.Boards{
				{
					Name:   "Magical board",
					Closed: false,
				},
				{
					Name:   "Adventure board",
					Closed: true,
				},
			}

			client.
				EXPECT().
				GetBoards().
				Return(boards, tc.err).
				AnyTimes()

			var b bytes.Buffer
			writer := bufio.NewWriter(&b)

			opts := board.ListOptions{
				ShowClosed: tc.closed,
			}

			board.DisplayBoards(client, opts, writer)
			writer.Flush()

			if b.String() != tc.output {
				t.Fatalf("expected '%s'; got '%s'", tc.output, b.String())
			}
		})
	}
}
