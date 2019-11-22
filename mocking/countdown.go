package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!"
const countDownStart = 3

// Countdown function prints out the countdown on the commandline
func Countdown(w io.Writer) {
	for i := countDownStart; i >= 1; i-- {
		fmt.Fprintln(w, i)
	}
	fmt.Fprintf(w, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
