package main

import (
	fm "fmt"
	"os"
)

// HappyNewYear is a long string
const HappyNewYear = "Happy new year!  Happy new year! Happy new year!" +
	"Happy new year! Happy new year!"

// MultiLineHappyNewYear is a multi line long string
const MultiLineHappyNewYear = `Happy new year! Happy new year! Happy new year! Happy new year!
 Happy new year! Happy new year! Happy new year! `

var v = 5

// T type definition demo
type T struct{}

func init() {
}

func main() {
	fm.Println("Hello, world")
	beyondHello()
	printLongString(HappyNewYear)
	printLongString(MultiLineHappyNewYear)

	// get os
	shell := os.Getenv("SHELL")
	fm.Printf("Current using shell is %s\n", shell)
	// get system path
	path := os.Getenv("PATH")
	fm.Printf("The system path is %s\n", path)

	printVariable()
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

func printVariable() {
	fm.Println(v)
	defAndPrintNewV()
}

func defAndPrintNewV() {
	v := 3
	fm.Println(v)
	v2()
}

func v2() {
	fm.Println(v)
}
