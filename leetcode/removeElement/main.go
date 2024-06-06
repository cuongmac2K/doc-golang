package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(arr, 2))
}
func removeElement(nums []int, val int) int {
	count := len(nums)
	for i, v := range nums {
		if v == val {
			nums[i] = 101
			count--
		}
	}

	sort.Ints(nums)
	fmt.Println(nums)
	return count
}
