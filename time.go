package main

import (
	"fmt"
	"time"
)

func timeSub() {
	start := time.Now()
	longCalculation()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)
}

func longCalculation() {
	for i := 1; i < 10000000; i++ {
		j := 0
		j++
		j--
	}
}
