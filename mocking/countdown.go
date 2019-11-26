package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countDownStart = 3

// Sleeper interface which represents our dependency of Sleep
type Sleeper interface {
	Sleep()
}

// Countdown function prints out the countdown on the commandline
func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countDownStart; i >= 1; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(w, finalWord)
}

type defaultSleeper struct{}

func (d *defaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	Countdown(os.Stdout, &defaultSleeper{})
}
