package main

import (
	"fmt"
)

// []int{4, 2, 3}
func missingNumbers(arr []int) []int {
	n := len(arr) + 2
	fmt.Printf("n : %d\n", n)
	expectedSum := n * (n + 1) / 2
	fmt.Printf("expectedSum : %d\n", expectedSum)
	actualSum := 0

	for _, num := range arr {
		actualSum += num
	}

	fmt.Printf("actualSum : %d\n", actualSum)

	missingSum := expectedSum - actualSum

	fmt.Printf("missingSum : %d\n", missingSum)

	// Find the midpoint of the missing sum
	midpoint := missingSum / 2

	fmt.Printf("midpoint : %d\n", midpoint)

	// Find the sum of numbers less than or equal to the midpoint
	sumSmaller := 0
	for _, num := range arr {
		if num <= midpoint {
			sumSmaller += num
		}
	}

	fmt.Printf("sumSmaller : %d\n", sumSmaller)

	// The difference between the sum of numbers less than or equal to the midpoint
	// and the expected sum of those numbers will give us one of the missing numbers
	missing1 := (midpoint * (midpoint + 1) / 2) - sumSmaller

	fmt.Printf("missing1 : %d\n", missing1)

	// The second missing number is the difference between the total missing sum
	// and the first missing number
	missing2 := missingSum - missing1

	fmt.Printf("missing2 : %d\n", missing2)

	return []int{missing1, missing2}
}

func main() {
	fmt.Println(missingNumbers([]int{4, 2, 3}))    // Output: [1, 5]
	fmt.Println(missingNumbers([]int{1, 2, 3, 4})) // Output: [5, 6]
	fmt.Println(missingNumbers([]int{1, 6, 4, 5}))
}
