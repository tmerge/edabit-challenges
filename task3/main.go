package main

import "fmt"

/* TASK: Create a function that repeats each character in a string n times. */

// printRepeatedString prints every character of input string <iterations> times.
func printRepeatedString(input string, iterations int) {
	var outputArray []byte

	// iterate through all letters and append it to byte array
	for i := 0; i < len(input); i++ {
		for j := 0; j < iterations; j++ {
			outputArray = append(outputArray, input[i])
		}
	}
	fmt.Printf("Repeated String: %s\n", string(outputArray))
}

func main() {

	// Test printRepeatedString with Hello World! and four iterations
	printRepeatedString("Hello World!", 4)

	// Test printRepeatedString with Peter and two iterations
	printRepeatedString("Peter", 2)
}
