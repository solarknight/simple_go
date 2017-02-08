package main

import (
	fm "fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
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
	fm.Println("Current system is", runtime.GOOS)
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

	// printVariable()
	// testRandom()
	// capsStr("Hello, world!")
	// switchDemo()
	// forRange()
	// funcDeclare()
	// deferFunc()
	// fm.Println(fibo(5))
	// funcParameter()
	// timeSub()
	// trySlice()
	// tryMap()
	// tryRegexp()
	// tryStruct()
	// tryInterface()
	// tryReflect()
	// tryScan()
	// tryFile()
	// tryJson()
	// tryGob()
	// tryError()
	// tryPanic()
	// tryExec()
	// tryCoroutine()
	// tryChannel()
	// trySelect()
	// tryLazyEvaluation()
	// tryDial()
	// trySocket()
	tryClosure()
	// tryHTTPReq()

	// start web server
	// startWebServer()
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

func testRandom() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fm.Printf("%d / ", a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fm.Printf("%d / ", r)
	}
	fm.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fm.Printf("%2.2f / \n", 100*rand.Float32())
	}
}
