package main

import (
	"fmt"
)

func missingNumbers(arr []int) []int {
	n := len(arr) + 2
	expectedSum := n * (n + 1) / 2
	actualSum := 0

	for _, num := range arr {
		actualSum += num
	}

	missingSum := expectedSum - actualSum

	// Find the midpoint of the missing sum
	midpoint := missingSum / 2

	// Find the sum of numbers less than or equal to the midpoint
	sumSmaller := 0
	for _, num := range arr {
		if num <= midpoint {
			sumSmaller += num
		}
	}

	// The difference between the sum of numbers less than or equal to the midpoint
	// and the expected sum of those numbers will give us one of the missing numbers
	missing1 := (midpoint * (midpoint + 1) / 2) - sumSmaller

	// The second missing number is the difference between the total missing sum
	// and the first missing number
	missing2 := missingSum - missing1

	return []int{missing1, missing2}
}

func main() {
	fmt.Println(missingNumbers([]int{4, 2, 3}))    // Output: [1, 5]
	fmt.Println(missingNumbers([]int{1, 2, 3, 4})) // Output: [5, 6]
	fmt.Println(missingNumbers([]int{1, 4, 5, 3}))
}
