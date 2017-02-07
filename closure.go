package main

import "fmt"

type node struct {
	data int
	next *node
}

func tryClosure() {
	tryCounter()
	tryPointer()
}

func tryCounter() {
	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())
}

func newCounter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func tryPointer() {
	a := new(node)
	fmt.Println("node", *a)
	update := nodeUpdate(a)
	update()
	fmt.Println("node", *a)
	update()
	fmt.Println("node", *a)
}

func nodeUpdate(a *node) func() {
	return func() {
		a.data = a.data + 1
	}
}
