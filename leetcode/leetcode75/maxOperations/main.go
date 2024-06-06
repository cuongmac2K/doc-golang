package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 1, 3, 4, 3}
	k := 6
	fmt.Println(maxOperations(a, k))
}
func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	i, j, count := 0, len(nums)-1, 0

	for i < j {
		sum := nums[i] + nums[j]
		if sum == k {
			count++
			i++
			j--
		} else if sum < k {
			i++
		} else {
			j--
		}
	}
	return count
}
