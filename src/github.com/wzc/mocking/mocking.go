package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const write = "write"
const sleep = "sleep"

//Sleeper 实现Sleep的都是Sleeper
type Sleeper interface {
	Sleep()
}

//DefaultSleeper uses in main
type DefaultSleeper struct{}

//SpySleeper make a mocking of sleep
type SpySleeper struct {
	Calls int
}

//CountdownOperationsSpy spy the order of write and sleep
type CountdownOperationsSpy struct {
	Calls []string
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

//Countdown is countdown
func Countdown(out io.Writer, sleeper Sleeper) {
	// //Fprint, Fprintf, Fprintln的区别
	// for i := countdownStart; i > 0; i-- {
	// 	sleeper.Sleep()
	// 	fmt.Fprintln(out, i)
	// }

	// sleeper.Sleep()
	// fmt.Fprint(out, "Go!")

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, "Go!")

	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
	}
	sleeper.Sleep()

}

//Sleep of SpySleeper is a func to count sleep times
func (s *SpySleeper) Sleep() {
	s.Calls++
}

//Sleep of DefaultSleeper
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

//Sleep of CountdownOperationsSpy
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

//Write of CountdownOperationsSpy
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
