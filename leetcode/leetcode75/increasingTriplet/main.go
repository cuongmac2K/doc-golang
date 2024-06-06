package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{5, 4, 7, 2, 1}
	//	a := []int{20, 100, 10, 12, 5, 13}
	fmt.Println(increasingTriplet(a))
}
func increasingTriplet(nums []int) bool {
	// Initialize two variables to track the smallest and second smallest elements
	first, second := math.MaxInt32, math.MaxInt32

	// Iterate through the elements in the input slice
	for _, num := range nums {
		// Check if the current number is smaller than or equal to the first smallest
		if num <= first {
			first = num
		} else if num <= second { // Check if the current number is smaller than or equal to the second smallest
			second = num
		} else { // If neither condition is met, an increasing triplet subsequence is found
			return true
		}
	}

	// No increasing triplet subsequence was found
	return false
}
