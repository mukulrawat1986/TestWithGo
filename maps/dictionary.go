package main

import "errors"

// Dictionary type based on the map[string]string type
type Dictionary map[string]string

// Search method to find the definition of a word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}

func main() {}
