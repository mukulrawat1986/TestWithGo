package main

import (
	"io"
	"os"
)

// Countdown function prints out the countdown on the commandline
func Countdown(w io.Writer) {}

func main() {
	Countdown(os.Stdout)
}
