package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// create a buffer where the printing happens so we can test it
	buffer := &bytes.Buffer{}

	Greet(buffer, "Chris")

	got := buffer.String()

	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
