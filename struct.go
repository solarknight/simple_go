package main

import (
	"fmt"
)

// User struct
type User struct {
	name string
	age  int
	sex  int
}

func tryStruct() {
	t := new(User)
	t.name = "Tom"
	t.age = 18
	fmt.Println(t)
}
