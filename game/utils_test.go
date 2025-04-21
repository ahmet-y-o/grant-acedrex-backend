package game_test

import (
	"acedrex/game"
	"testing"
)

func TestNotationToCoords(t *testing.T) {
	result, err := game.CoordsToNotation(game.Coords{0, 0})
	want := "a12"
	if err != nil {
		t.Fatalf("CoordsToNotation(Coords{0, 0}) = %v", err)
	}
	if result != want {
		t.Fatalf("CoordsToNotation(Coords{0, 0}) = %v, want %v", result, want)
	} else {
		t.Logf("CoordsToNotation(Coords{0, 0}) = %v", result)
	}
}
