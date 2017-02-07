package main

import (
	"fmt"
	"time"
)

func tryChannel() {
	fmt.Println()
	// basicChannelUsage()
	closeChannel()
}

func basicChannelUsage() {
	ch1 := make(chan string)
	ch2 := make(chan int)
	defer close(ch1)
	defer close(ch2)

	go func() {
		ch1 <- "Hello, world!"
	}()
	go func() {
		ch2 <- 1
	}()

	go func() {
		if v, ok := <-ch2; v == 1 && ok {
			fmt.Println("From channel get string:", <-ch1)
		}
	}()
	time.Sleep(1e9)
}

func closeChannel() {
	ch := make(chan string)
	go sendData(ch)
	// getData(ch)
	getData2(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}

func getData2(ch chan string) {
	for input := range ch {
		fmt.Printf("%s ", input)
	}
}
