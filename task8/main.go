package main

import (
	"fmt"
	"regexp"
)

/* TASK: Create a function that takes a string, transforms all but the last four characters into "#" and returns the new masked string.*/

// maskifyString takes an string and replace all numbers til the last four to '#'
func maskifyString(input string) {
	inputAsArray := []byte(input)
	// Check input for expression and process if its true
	matched, err := regexp.Match("[0-9]*", inputAsArray)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	if matched {
		// Handle special case
		if input == "" {
			fmt.Print("Masked String: <empty>\n")
			return
		}
		// Convert chars to '#'
		for i := 0; i < len(input)-4; i++ {
			inputAsArray[i] = '#'
		}
		fmt.Print("Masked String: " + string(inputAsArray) + "\n")
	}
	return
}

func main() {

	// Test maskifyString with 12345678 -> Expect: ####5678
	maskifyString("12345678")

	// Test maskifyString with empty string -> Expect: <empty>
	maskifyString("")

	// Test maskifyString with 4556364607935616 -> Expect: ############5616
	maskifyString("4556364607935616")

	// Test maskifyString with 1 -> Expect: 1
	maskifyString("1")
}
