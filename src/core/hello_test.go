package core

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "Hello Me"
	resp := Hello("Me")
	if resp != expected {
		t.Errorf("received %v; expected %v", resp, expected)
	}
}
