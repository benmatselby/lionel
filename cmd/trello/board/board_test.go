package board_test

import (
	"testing"

	"github.com/benmatselby/lionel/cmd/trello/board"
	"github.com/benmatselby/lionel/mock_trello"

	"github.com/golang/mock/gomock"
)

func TestNewBoardCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mock_trello.NewMockAPI(ctrl)

	cmd := board.NewBoardCommand(client)

	use := "board"
	short := "Commands that relate to Trello boards"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}
