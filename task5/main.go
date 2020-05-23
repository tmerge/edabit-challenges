package main

import (
	"fmt"
	"regexp"
)

/* TASK: Create a function that determines whether a string is a valid hex code.
A hex code must begin with a pound key # and is exactly 6 characters in length.
Each character must be a digit from 0-9 or an alphabetic character from A-F.
All alphabetic characters may be uppercase or lowercase.
*/

// isValidHexCode take a string and check if its provides pattern
func isValidHexCode(input string) bool {

	// apply regex expression to byte stream
	matched, err := regexp.Match(`^#[0-9a-fA-F]{6}`, []byte(input))
	if err != nil {
		fmt.Print(err.Error())
		return false
	}
	return matched
}

// TODO: find misstake in regex expression?
func main() {

	// Test isValidHexCode with #6ADEEF -> Expect: true
	fmt.Printf("Result: %t\n", isValidHexCode("#6ADEEF"))

	// Test isValidHexCode with 6ADEEF -> Expect: false
	fmt.Printf("Result: %t\n", isValidHexCode("6ADEEF"))

	// Test isValidHexCode with #Z6ADEE -> Expect: false
	fmt.Printf("Result: %t\n", isValidHexCode("#Z6ADEE"))
}
