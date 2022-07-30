package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countDownStart = 3
	write          = "write"
	sleep          = "sleep"
)

type Sleeper interface {
	Sleep()
}

type spyCountDownOperations struct {
	Calls []string
}

func (s *spyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *spyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleep struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleep) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleep{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
