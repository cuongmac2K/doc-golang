package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 4
	fmt.Println(searchInsert(nums, target))
}
func searchInsert(nums []int, target int) int {
	for i, d := range nums {
		if d == target {
			return i
		} else if d > target {
			return i
		}
	}
	return len(nums)
}
