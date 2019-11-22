package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown function prints out the countdown on the commandline
func Countdown(w io.Writer) {
	fmt.Fprintf(w, "3")
}

func main() {
	Countdown(os.Stdout)
}
