package main

import fm "fmt"

// HappyNewYear is a long string
const HappyNewYear = "Happy new year!  Happy new year! Happy new year!" +
	"Happy new year! Happy new year!"

// MultiLineHappyNewYear is a multi line long string
const MultiLineHappyNewYear = `Happy new year! Happy new year! Happy new year! Happy new year!
 Happy new year! Happy new year! Happy new year! `

func main() {
	fm.Println("Hello, world")
	beyondHello()
	printLongString(HappyNewYear)
	printLongString(MultiLineHappyNewYear)
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

func printLongString(s string) {
	fm.Println(s)
}
