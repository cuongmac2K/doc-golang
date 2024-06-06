package main

import (
	"fmt"
	"sort"
)

func main() {
	//a := []int{-5, 5, 4, -3, 0, 0, 4, -2}
	a := []int{1, 0, -1, 0, -2, 2}
	//a := []int{2, 2, 2, 2, 2}
	fmt.Println(fourSum(a, 0))
}
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, target, 0, 4)
}
func kSum(nums []int, target, start, k int) [][]int {
	var res [][]int
	if start == len(nums) {
		return res
	}
	avgValue := target / k

	if nums[start] > avgValue || avgValue > nums[len(nums)-1] {
		return res
	}
	if k == 2 {
		return twoSum(nums, target, start)
	}
	for i := start; i < len(nums); i++ {
		if i == start || nums[i-1] != nums[i] {
			for _, subset := range kSum(nums, target-nums[i], i+1, k-1) {
				res = append(res, append([]int{nums[i]}, subset...))
			}
		}
	}
	return res
}
func twoSum(nums []int, target, start int) [][]int {
	res := make([][]int, 0)
	low, high := start, len(nums)-1
	for low < high {
		currSum := nums[low] + nums[high]
		if currSum < target || (low > start && nums[low] == nums[low-1]) {
			low++
		} else if currSum > target || high < len(nums)-1 && nums[high] == nums[high+1] {
			high--
		} else {
			res = append(res, []int{nums[low], nums[high]})
			low++
			high--
		}
	}
	return res
}
