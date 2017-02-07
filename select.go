package main

import (
	"fmt"
	"time"
)

func trySelect() {
	// tryListen()
	// tryTicker()
	tryTickerAndTimer()
}

func tryListen() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(0.01 * 1e9)
}

func pump1(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 <-chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received value from channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received value from channel 2: %d\n", v)
		}
	}

}

func tryTicker() {
	var dur time.Duration = 1e8
	chRate := time.Tick(dur) // a tick every 1/10th of a second
	timeCall := timePrint()

	i := 0
	for {
		<-chRate // rate limit our Service.Method RPC calls
		if i++; i > 10 {
			break
		}
		go timeCall()
	}
}

func timePrint() func() {
	i := 0
	return func() {
		i++
		fmt.Println("Timer call ", i)
	}
}

func tryTickerAndTimer() {
	var dur time.Duration = 1e8
	chRate := time.Tick(dur) // a tick every 1/10th of a second
	chTimer := time.After(time.Second)
	timeCall := timePrint()

	for {
		select {
		case <-chRate:
			timeCall()
		case <-chTimer:
			fmt.Println("should break now!")
			return
		}
	}
}
