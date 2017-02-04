package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func tryScan() {
	fmt.Println("Please enter your full name: ")
	// fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	// fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
	// 输出结果: From the string we read: 56.12 5212 Go

	// inputReader := bufio.NewReader(os.Stdin)
	// fmt.Println("Please enter some input: ")
	// input, err := inputReader.ReadString('\n')
	// if err == nil {
	// fmt.Printf("The input was: %s\n", input)
	// }
}
