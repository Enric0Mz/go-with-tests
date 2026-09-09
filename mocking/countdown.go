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
type SpyCoundDownOperations struct {
	Calls []string
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpyCoundDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCoundDownOperations) Write(p []byte) (i int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const countDonwStart = 3
const finalLine = "Go!"

const write = "write"
const sleep = "sleep"

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
