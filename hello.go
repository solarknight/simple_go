package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, world\n")
		beyondHello()
}

func beyondHello() {
		var x int = 3
		y := 4
		sum, prod := learnMultiple(x, y)
		fmt.Println("sum:", sum , "prod:", prod)
}

func learnMultiple(x, y int) (sum, prod int) {
		return x + y, x * y
}

