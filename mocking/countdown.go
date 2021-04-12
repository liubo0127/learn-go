package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownStart = 3
	finalWorld = "Go!"
)

type Sleeper interface {
	Sleep()
}

type SpySleep struct {
	Calls int
}

func (s *SpySleep) Sleep() {
	s.Calls++
}

type ConfigSleeper struct {
	duration time.Duration
}

func (c *ConfigSleeper) Sleep() {
	time.Sleep(c.duration)
}

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintln(out, finalWorld)
}

func main() {
	sleeper := &ConfigSleeper{time.Second}
	CountDown(os.Stdout, sleeper)
}
