package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countDownStart = 3

// Countdown function prints out the countdown on the commandline
func Countdown(w io.Writer) {
	for i := countDownStart; i >= 1; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(w, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprintf(w, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
