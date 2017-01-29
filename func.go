package main

import "fmt"

type binOp func(int, int) int

var a = 5
var b = 6

func funcDeclare() {
	fmt.Printf("add %d and %d result %d\n", a, b, unNamedAdd(a, b))
	fmt.Printf("add %d and %d result %d\n", a, b, namedAdd(a, b))

	changeValue(&a, 3)
	changeValue(&b, 4)
	fmt.Printf("add %d and %d result %d\n", a, b, unNamedAdd(a, b))
	fmt.Printf("add %d and %d result %d\n", a, b, namedAdd(a, b))
	changeableParam(8, 2, 5, 3, 6)
}

func unNamedAdd(a, b int) int {
	return a + b
}

func namedAdd(a, b int) (result int) {
	result = a + b
	return
}

func changeValue(target *int, value int) {
	*target = value
}

func changeableParam(no1 int, arg ...int) {
	fmt.Print(no1)
	for _, v := range arg {
		fmt.Print(v)
	}
}
