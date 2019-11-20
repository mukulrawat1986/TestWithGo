package main

// Dictionary type based on the map[string]string type
type Dictionary map[string]string

// Search method to find the definition of a word in the dictionary
func (d Dictionary) Search(word string) string {
	return d[word]
}

func main() {}
