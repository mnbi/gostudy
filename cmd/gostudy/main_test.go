package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	msg := greeting()
	if msg != "Hello, world!" {
		t.Fatalf(`greeting() = %v, want "Hello, world!"`, msg)
	}
}
