package main

import "fmt"

func main() {
	//a := []int{1, 2, 0}
	a := []int{3, 4, -1, 1}
	//a := []int{7, 8, 9, 11, 12}
	fmt.Println(firstMissingPositive(a))
}
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] >= 1 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i-- // keep the same index
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
