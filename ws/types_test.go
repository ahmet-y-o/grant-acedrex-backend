package ws

import (
	"testing"
)

func TestTokenizeMoveMessage(t *testing.T) {
	msg := "a11b7"
	want := []string{"a", "11", "b", "7"}
	got, err := TokenizeMoveMessage(msg)
	if err != nil {
		t.Fatalf("TokenizeMoveMessage(%q) = %v", msg, err)
	}
	if len(want) != len(got) {
		t.Fatalf("TokenizeMoveMessage(%q) = %v, want %v", msg, got, want)
	}
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("TokenizeMoveMessage(%q) = %v, want %v", msg, got, want)
		}
	}
}
