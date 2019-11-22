package main

import (
	"fmt"
	"io"
	"os"
)

// Greet function greets a person
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elode")
}
