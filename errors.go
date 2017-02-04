package main

import "fmt"

func tryError() {
	var ret int
	var err error
	if ret, err = printNumber(-3); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println(ret)
}

func printNumber(a int) (int, error) {
	if a < 0 {
		// return a, errors.New("no negative number")
		return a, fmt.Errorf("negative number %d not accepted", a)
	}
	fmt.Println("Received number is", a)
	return a, nil
}
