package main

import (
	"fmt"
	"io"
)

// Greet function greets a person
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func main() {}
