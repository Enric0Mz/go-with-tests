package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const countDonwStart = 3
const finalLine = "Go!"

func Countdown(writter io.Writer, sleeper Sleeper) {

	for i := countDonwStart; i > 0; i-- {
		fmt.Fprintln(writter, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writter, finalLine)

}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
