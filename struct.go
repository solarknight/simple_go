package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// User struct
type User struct {
	name string
	age  int
	sex  int
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

type number struct {
	f float32
}

type nr number // alias type

// File struct
type File struct {
	fd   int    // 文件描述符
	name string // 文件名
}

// TagType struct
type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS //anonymous field
}

type A struct{ a int }
type B struct{ a, b int }

type C struct {
	A
	B
}

type TwoInts struct {
	a int
	b int
}

func tryStruct() {
	t := new(User)
	t.name = "Tom"
	t.age = 18
	fmt.Println(t)

	p := &User{"Cherry", 15, 1}
	fmt.Println(p)

	a := &number{5.0}
	b := &nr{5.0}
	var c = number(*b)
	fmt.Println(a, b, c)

	file := NewFile(10, "network")
	size := unsafe.Sizeof(file)
	fmt.Println("file size is", size)

	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}

	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)

	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())
	fmt.Printf("The prod is %d\n", two2.prod())
}

// NewFile method
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

func (tn *TwoInts) prod() int {
	return tn.a * tn.b
}
