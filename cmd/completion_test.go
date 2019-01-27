package cmd_test

import (
	"testing"

	"github.com/benmatselby/lionel/cmd"
	"github.com/spf13/cobra"
)

func TestNewCompletionCommand(t *testing.T) {
	root := &cobra.Command{
		Use: "mock",
	}

	cmd := cmd.NewCompletionCommand(root)

	use := "completion"
	short := "Generates the bash completion script: lionel.sh"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}
