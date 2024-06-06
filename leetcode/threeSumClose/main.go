package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{-1, 2, 1, -4}
	target := 3
	fmt.Println(threeSumClosest(a, target))
}
func threeSumClosest(nums []int, target int) int {

	// Helper func
	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}

	// Sort the input array in ascending order
	sort.Ints(nums)

	// Initialize the result variable with the sum of the first three elements
	result := nums[0] + nums[1] + nums[2]

	// Iterate through the array
	for i := 0; i < len(nums)-2; i++ {

		// Use two pointers approach
		left := i + 1
		right := len(nums) - 1

		// Continue until the two pointers meet
		for left < right {
			// Calculate the sum of three elements
			sum := nums[i] + nums[left] + nums[right]

			// If the sum is equal to the target, return the sum
			if sum == target {
				return sum
			}

			// Update the result if the current sum is closer to the target
			if abs(target-sum) < abs(target-result) {
				result = sum
			}

			// Move the pointers based on the sum
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
