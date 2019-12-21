package main

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "https://www.facebook.com"
	fastURL := "https://www.google.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
