package main

import (
	"fmt"
)

/* TASK: Create a function that takes a strings characters as ASCII and returns each characters hexadecimal value as a string. */

// ConvertToHex take a string and convert it to hex values then print it
func ConvertToHex(input string) {
	inputAsArray := []byte(input)

	// Print hex values with whitespaces
	fmt.Print(input + ": ")
	for i := 0; i < len(inputAsArray); i++ {
		fmt.Printf("%X ", inputAsArray[i])
	}
	// Print break line to prettfy output
	fmt.Print("\n")
}

func main() {

	// Test ConvertToHex with "hello world" -> Expect: 68 65 6c 6c 6f 20 77 6f 72 6c 64
	ConvertToHex("hello world")

	// Test ConvertToHex with "Welcome" -> Expect: 57 65 6C 63 6F 6D 65
	ConvertToHex("Welcome")

}
