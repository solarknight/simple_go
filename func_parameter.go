package main

import "fmt"

func funcParameter() {
	fplus := func(x, y int) int { return x + y }
	fmt.Println(fplus(2, 3))
	func(x, y int) int { return x + y }(2, 3)
}
