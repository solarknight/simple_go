package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var str = "hello"
var map1 = map[string]int{"one": 1, "two": 2, "three": 3}

type robot struct {
	name  string
	power int
}

type human struct {
	name  string
	power int
}

type worker interface {
	work() int
}

func (a *robot) work() int {
	fmt.Println("I'm working")
	return a.power
}

func (a *human) work() int {
	fmt.Println("I'm working")
	return a.power
}

func practise() {
	// stringPractise()
	// mapPractise()
	// structPractise()
	interfacePractise()
}

func stringPractise() {
	modifyString()
	subString()
	iterateString()
	stringLen()
	stringJoin()
}

func modifyString() {
	c := []byte(str)
	c[0] = 'w'
	fmt.Println(string(c))
}

func subString() {
	s1 := str[2:4]
	fmt.Println(s1)
	s2 := str[1:]
	fmt.Println(s2)
}

func iterateString() {
	for ix, c := range str {
		fmt.Printf("idx: %d, char: %c\n", ix, c)
	}
}

func stringLen() {
	fmt.Println("str len is", len(str))
	fmt.Println("str len is", utf8.RuneCountInString(str))
}

func stringJoin() {
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined result %s\n", str3)

	str4 := " world"
	str4 = str + str4
	fmt.Println(str4)
}

func mapPractise() {
	for ix, v := range map1 {
		fmt.Println("idx:", ix, "value:", v)
	}

	key1 := "two"
	if v, present := map1[key1]; present {
		fmt.Println(v)
	}

	delete(map1, key1)
	if v, present := map1[key1]; present {
		fmt.Println(v)
	} else {
		fmt.Println("key does not exist")
	}
}

func structPractise() {
	robot1 := new(robot)
	robot1.name = "Kfc443"
	robot1.power = 5
	fmt.Println(robot1)

	robot2 := &robot{"Mlc80", 999}
	fmt.Println(robot2)

	robot3 := newRobot("Ll88", 777)
	fmt.Println(robot3)
}

func newRobot(name string, power int) *robot {
	return &robot{name, power}
}

func interfacePractise() {
	var worker1 worker
	a := newRobot("Tom", 20)
	worker1 = a

	if _, ok := worker1.(*robot); ok {
		fmt.Println("worker is a robot!")
	}

	b := &human{"Lilei", 5}
	worker1 = b
	switch worker1.(type) {
	case *robot:
		fmt.Println("worker is a robot!")
	case *human:
		fmt.Println("worker is a human!")
	default:
		fmt.Println("worker is unknown!")
	}
}
