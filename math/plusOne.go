package main

import "fmt"

// plusOne increments an integer represented as a slice of digits.
// The function handles carry-over correctly (e.g., 9 → 0 and carry).
// It only allocates new memory when the input consists entirely of 9's.

func plusOne(digits []int) []int {
	n := len(digits) - 1

	// Process digits from right to left
	for i := n; i >= 0; i-- {

		// If digit is less than 9, simply increment and return
		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		// Otherwise, digit is 9 → becomes 0 and carry continues
		digits[i] = 0
	}

	// If all digits were 9, prepend 1 at the front (e.g., 999 → 1000)
	return append([]int{1}, digits...)
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3})) // [1 2 4]
	fmt.Println(plusOne([]int{9, 9, 9})) // [1 0 0 0]
}
