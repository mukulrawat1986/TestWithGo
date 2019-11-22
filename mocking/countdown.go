package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown function prints out the countdown on the commandline
func Countdown(w io.Writer) {
	for i := 3; i >= 1; i-- {
		fmt.Fprintln(w, i)
	}
	fmt.Fprintf(w, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
