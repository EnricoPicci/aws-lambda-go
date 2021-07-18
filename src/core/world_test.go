package core

import (
	"testing"
)

func TestWorld(t *testing.T) {
	expected := "The world is waiting for you, You"
	resp := World("You")
	if resp != expected {
		t.Errorf("received %v; expected %v", resp, expected)
	}
}
