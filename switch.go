package main

import (
	"fmt"
)

func switchDemo() {
	num1 := 99
	switch num1 {
	case 99:
		fmt.Println("hit 99!")
		fallthrough
	case 100:
		fmt.Println("hit 100!")
	default:
		fmt.Println("no number hited!")
	}
}
