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

// Node struct
type Node struct {
	data float64
	next *Node
}

// DNode double linked node
type DNode struct {
	head *DNode
	data int
	tail *DNode
}

// Tree tree node
type Tree struct {
	le   *Tree
	data float64
	ri   *Tree
}

func tryStruct() {
	t := new(User)
	t.name = "Tom"
	t.age = 18
	fmt.Println(t)

	p := &User{"Cherry", 15, 1}
	fmt.Println(p)
}
