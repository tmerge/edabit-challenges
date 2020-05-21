package main

import (
	"fmt"
	"strings"
)

/* TASK: Create a function that returns true if an input string contains only uppercase or only lowercase letters. */

// checkForCase returns a boolean expression, if input contains only lower or uppercase letters it will be true, else flase.
func checkForCase(input string) bool {

	compareStringUpper := strings.ToUpper(input)
	compareStringLower := strings.ToLower(input)

	// compare input with compareStrings
	if input == compareStringUpper {
		return true
	}

	if input == compareStringLower {
		return true
	}

	return false
}

func main() {

	// Result should be true for sucess
	testUpper := "HIHI"
	onlyUpp := checkForCase(testUpper)
	fmt.Printf("onlyUpp Test: %t\n", onlyUpp)

	// Result should be true for sucess
	testLower := "hihi"
	onlyLow := checkForCase(testLower)
	fmt.Printf("onlyLow Test: %t\n", onlyLow)

	// Result should be false for sucess
	testMixed := "HiHi"
	mixed := checkForCase(testMixed)
	fmt.Printf("mixed Test: %t\n", mixed)
}
