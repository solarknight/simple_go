package main

import "fmt"
import "sort"

var slice1 = array1[1:4]
var slice2 = &array2
var slice3 = array3[:]
var slice4 = array4[0:]
var slice5 = array5[:len(array5)-1]

var slice6 = make([]int, 10)
var slice7 = make([]int, 4, 8)
var slice8 = new([100]int)[0:50]

func trySlice() {
	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Printf("%s\n", s1)
	fmt.Printf("%s\n", s2)

	s2[1] = 't'
	fmt.Printf("%s\n", s1)
	fmt.Printf("%s\n", s2)

	c1From := []int{1, 2, 3}
	c1To := make([]int, 10)
	copy(c1To, c1From)
	fmt.Println(c1To)
	c3 := []int{1, 3, 2}
	c3 = append(c3, 4, 5, 6)
	fmt.Println(c3)

	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s3 := string(c) // s2 == "cello"
	fmt.Println(s)
	fmt.Println(s3)

	sliceSort()
}

func sliceSort() {
	sort.Ints(slice1)
	fmt.Println(slice1)
}
