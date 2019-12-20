package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"https://www.google.com":   true,
		"https://www.facebook.com": true,
		"waat://furhurterwe.geds":  false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
