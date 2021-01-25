package poker_test

import (
	"bytes"
	"strings"
	"testing"

	poker "github.com/nanduzz/learn-go-with-tests/application"
)

func TestGame(t *testing.T) {
	stdout := &bytes.Buffer{}
	in := strings.NewReader("7\n")
	game := &poker.GameSpy{}

	cli := poker.NewCLI(in, stdout, game)
	cli.PlayPoker()

	gotPrompt := stdout.String()
	wantPrompt := poker.PlayerPrompt

	if gotPrompt != wantPrompt {
		t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
	}

	if game.StartedWith != 7 {
		t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
	}
}
