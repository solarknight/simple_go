package main

import (
	"fmt"
	"strconv"
	"strings"
)

func capsStr(str string) {
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))
	if strings.Contains(str, "Hello") {
		fmt.Println("Str", str, "index", strings.Index(str, "Hello"))
		fmt.Println(strings.Count(str, "o"))
		fmt.Println(strings.Replace(str, "o", "e", -1))
	}
	sArray := strings.Fields(str)
	fmt.Println(strings.Join(sArray, "&&"))
	sArray2 := strings.Split(str, "o")
	fmt.Println(strings.Join(sArray2, "&&"))

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
	// var intP *int
	// intP = &i1
	intP := &i1
	fmt.Printf("Pointer location is %p, value is %d\n", intP, *intP)
}
