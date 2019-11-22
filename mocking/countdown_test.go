package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	// we create a buffer where the printing takes place that
	// we can test
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
