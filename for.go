package main

import "fmt"

func forRange() {
	str := "Hello, world!"
	for idx, s := range str {
		fmt.Printf("idx: %d, char: %c\n", idx, s)
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
}
