package main

import "errors"

var (
	// ErrNotFound is the error returned when a word does not exist in the
	// dictionary
	ErrNotFound = errors.New("could not find the word you were looking for")
)

// Dictionary type based on the map[string]string type
type Dictionary map[string]string

// Search method to find the definition of a word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func main() {}
