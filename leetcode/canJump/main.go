package main

import "fmt"

func main() {
	//a := []int{2, 3, 1, 1, 4}
	a := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(a))
}
func canJump(nums []int) bool {
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if prev == 0 {
			return false
		}
		prev = max(prev-1, nums[i])
	}
	return true
}
