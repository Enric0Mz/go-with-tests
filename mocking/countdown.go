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

type SpyCoundDownOperations struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *SpyCoundDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
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
	secondsToSleep := 2 * time.Second
	sleeper := &ConfigurableSleeper{secondsToSleep, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
