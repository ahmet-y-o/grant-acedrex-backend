package game_test

import (
	"acedrex/game"
	"testing"
)

func TestGAFEN(t *testing.T) {
	g := game.InitilaizeGame()
	want := "rlugcakcgulr/12/12/pppppppppppp/12/12/12/12/PPPPPPPPPPPP/12/12/RLUGCAKCGULR w Jj"
	got := g.GAFENotation()
	if want != got {
		t.Fatalf("\nGAFENotation() = %v,\n want %v", got, want)
	}
}
