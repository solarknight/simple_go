package main

import fm "fmt"

func main() {
	fm.Println("Hello, world")
	beyondHello()
}

func beyondHello() {
	var x = 2
	y := 7
	sum, prod := learnMultiple(x, y)
	fm.Println("sum:", sum, "prod:", prod)
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}
