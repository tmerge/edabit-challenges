package main

import (
	"fmt"
	"regexp"
)

/* TASK: ATM machines allow 4 or 6 digit PIN codes and PIN codes cannot contain anything but exactly 4 digits or exactly 6 digits.
Your task is to create a function that takes a string and returns true if the PIN is valid and false if it's not.
*/

// validatePin takes an string and check it for expression matches and return bool
func validatePin(input string) bool {
	matched, err := regexp.Match("^[0-9]{4,6}", []byte(input))
	if err != nil {
		fmt.Print(err.Error())
		return false
	}
	return matched
}

func main() {

	// Test validatePin with 1234 -> Expect: true
	fmt.Println(validatePin("1234"))

	// Test validatePin with 123456 -> Expect: true
	fmt.Println(validatePin("123456"))

	// Test validatePin with a234 -> Expect: false
	fmt.Println(validatePin("a234"))

	// Test validatePin with <empty> -> Expect: false
	fmt.Println(validatePin(""))
}
