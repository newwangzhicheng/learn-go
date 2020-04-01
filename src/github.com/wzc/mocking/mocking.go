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

//ConfigurableSleeper 可配置的sleeper
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

//SpyTime is struct iof duration
type SpyTime struct {
	durationSlept time.Duration
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)

	configurableSleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, configurableSleeper)
}

//Countdown is countdown
func Countdown(out io.Writer, sleeper Sleeper) {
	//Fprint, Fprintf, Fprintln的区别
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, "Go!")

	// for i := countdownStart; i > 0; i-- {
	// 	fmt.Fprintln(out, i)
	// }
	// fmt.Fprint(out, "Go!")

	// for i := countdownStart; i > 0; i-- {
	// 	sleeper.Sleep()
	// }
	// sleeper.Sleep()

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
//实现了Sleeper接口
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

//Write of CountdownOperationsSpy
//实现了io.Writer的接口
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

//Sleep of SpyTime
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

//Sleep of ConfigurableSleeper
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
