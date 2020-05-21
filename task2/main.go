package main

import (
	"fmt"
)

/* TASK: Create a function that takes a number as an argument.
Add up all the numbers from 1 to the number you passed to the function.
For example, if the input is 4 then your function should return 10 because 1 + 2 + 3 + 4 = 10.
*/

// addUp add positive numbers from one to size of number argument, for exmaple addUp(3) = 1 + 2 + 3 = 6
func addUp(number int) {
	result := 0

	// Check for numbers less than zero
	if number < 0 {
		fmt.Printf("Numbers less than zero are not allowed (Your number: %d).\n", number)
		return
	}

	for i := 1; i <= number; i++ {
		result += i
	}

	// Print result to terminal
	fmt.Printf("Result: %d\n", result)
}

func main() {

	// Test for addUp function with 3
	addUp(3)

	// Test for addUp function with -1
	addUp(-1)
}
