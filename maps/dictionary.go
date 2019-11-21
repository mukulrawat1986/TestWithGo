package main

const (
	// ErrNotFound is the error returned when a word does not exist in the
	// dictionary
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists is the error returned when we try to add a word which
	// already exists in the dictionary
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

// DictionaryErr type that implements the error interface
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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

// Add method to add new word and definition in the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update method updates the definition of a word present in the dictionary`
func (d Dictionary) Update(word, definition string) error {
	d[word] = definition
	return nil
}

func main() {}
