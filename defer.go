package main

import "fmt"

func deferFunc() {
	function1()
	defer2()
	defer3()
	bbb()
}

func function1() {
	fmt.Println("In function1 at the top")
	defer function2()
	fmt.Println("In function1 at the bottom!")
}

func function2() {
	fmt.Println("function2: Deferred until the end of the calling function!")
}

func defer2() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func defer3() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func aaa() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func bbb() {
	defer un(trace("b"))
	fmt.Println("in b")
	aaa()
}
