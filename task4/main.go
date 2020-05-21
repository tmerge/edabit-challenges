package main

import (
	"fmt"
)

/* TASK: Create a function that outputs true if a number is prime, and false otherwise. */

// isPrime takes a number and check number for prime and returns true if number is prime and false otherwise.
func isPrime(number int) bool {

	if number >= 1 {
		// one isn't prime and two is prime but fails modulo operator condition
		if number == 1 {
			return false
		}
		if number == 2 {
			return true
		}

		// check number with modulo operator for prime
		if number%2 == 0 {
			return false
		}
		return true
	}
	return false
}

func main() {

	// Test isPrime with 31
	fmt.Printf("Result: %t\n", isPrime(31))

	// Test isPrime with 18
	fmt.Printf("Result: %t\n", isPrime(18))

	// Test special cases with negativ number and 1
	fmt.Printf("Result: %t\n", isPrime(-1))
	fmt.Printf("Result: %t\n", isPrime(1))
}
