package main

import "fmt"

func tryPanic() {
	protect(badFunc, -1)
	recoverStack()
}

func protect(g func(a ...interface{}), a ...interface{}) {
	defer func() {
		fmt.Println("done")
		// Println executes normally even if there is a panic
		if err := recover(); err != nil {
			fmt.Printf("run time panic: %v\n", err)
		}
	}()
	g(a) //   possible runtime-error
}

func badFunc(a ...interface{}) {
	if len(a) > 0 {
		panic("A server error occured")
	}
}

func recoverStack() {
	f1()
	fmt.Println("Returned normally from f.")
}

func f1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
