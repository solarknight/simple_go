package main

import (
	"fmt"
	"regexp"
)

func tryRegexp() {
	fmt.Println()
	str1 := "Hello, world!"
	pattern := "Hel"
	matched, err := regexp.MatchString(pattern, str1)
	fmt.Println("matched:", matched, err)
}
