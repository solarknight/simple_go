package main

import (
	"fmt"
	"time"
)

type node struct {
	data int
	next *node
}

func tryClosure() {
	tryCounter()
	// tryPointer()
	// closureAndCoroutine()
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

var values = [5]int{10, 11, 12, 13, 14}

func closureAndCoroutine() {
	// 版本A:
	for ix := range values { // ix是索引值
		func() {
			fmt.Print(ix, " ")
		}() // 调用闭包打印每个索引值
	}
	fmt.Println()
	// 版本B: 和A版本类似，但是通过调用闭包作为一个协程
	for ix := range values {
		go func() {
			fmt.Print(ix, " ")
		}()
	}
	fmt.Println()
	time.Sleep(5e9)
	// 版本C: 正确的处理方式
	for ix := range values {
		go func(ix interface{}) {
			fmt.Print(ix, " ")
		}(ix)
	}
	fmt.Println()
	time.Sleep(5e9)
	// 版本D: 输出值:
	for ix := range values {
		val := values[ix]
		go func() {
			fmt.Print(val, " ")
		}()
	}
	fmt.Println()
	time.Sleep(1e9)
	// version E
	var val int
	for ix := range values {
		val = values[ix]
		go func() {
			fmt.Print(val, " ")
		}()
	}
	fmt.Println()
	time.Sleep(1e9)
}
