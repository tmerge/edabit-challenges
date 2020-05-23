package main

import (
	"fmt"
	"strconv"
)

/* TASK: Count the amount of ones in the binary representation of an integer.
So for example, since 12 is '1100' in binary, the return value should be 2.
*/

// countBinaryDigitOne takes a base int64 number and counts ones in base 2
func countBinaryDigitOne(number int64) int {
	binary := strconv.FormatInt(number, 2)
	counter := 0
	for i := 0; i < len(binary); i++ {
		if binary[i] == '1' {
			counter++
		}
	}
	return counter
}

func main() {

	// Test countBinaryDigitOne with 5 -> Expect: 2
	fmt.Printf("Count: %d\n", countBinaryDigitOne(5))

	// Test countBinaryDigitOne with 0 -> Expect: 0
	fmt.Printf("Count: %d\n", countBinaryDigitOne(0))

	// Test countBinaryDigitOne with 12 -> Expect: 2
	fmt.Printf("Count: %d\n", countBinaryDigitOne(12))

	// Test countBinaryDigitOne with 999 -> Expect: 8
	fmt.Printf("Count: %d\n", countBinaryDigitOne(999))

}
