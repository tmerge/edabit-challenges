package main

import (
	"fmt"
)

/* TASK: Create a function that takes two numbers as arguments (num, length) and returns an array of multiples of num up to length. */

// multiplyArrayOfNumber takes a number and a length and multiply number * number foreach step until length is reached
func multiplyArrayOfNumber(number int, length int) []int {
	slice := make([]int, length)
	for i := 0; i < len(slice); i++ {
		if i == 0 {
			slice[i] = number
			continue
		}
		// multiply index -1 by two and assign it to current index
		slice[i] = (slice[i-1] * 2)
	}
	return slice
}

func main() {

	// Test multiplyArrayOfNumber with 2 and 5 -> Expect 2,4,8,16,32
	array := multiplyArrayOfNumber(2, 5)
	fmt.Print("Result: ")
	for _, number := range array {
		fmt.Printf("%d ", number)
	}

}
