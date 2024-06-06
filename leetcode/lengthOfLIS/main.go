package main

import "fmt"
import "math"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS1(nums []int) int {
	return LISRecur(nums, math.MinInt32)
}

func LISRecur(nums []int, last int) int {
	if len(nums) == 0 {
		return 0
	}
	max := LISRecur(nums[1:], last)
	cur := 0
	if nums[0] > last {
		cur = 1 + LISRecur(nums[1:], nums[0])
	}
	if cur > max {
		max = cur
	}
	return max
}

func lengthOfLIS(nums []int) int {
	t := make([]int, len(nums))
	var s int
	for _, v := range nums {
		var i, j = 0, s
		for i != j {
			m := (i + j) / 2
			if t[m] < v {
				i = m + 1
			} else {
				j = m
			}
		}
		t[i] = v
		if i == s {
			s++
		}
	}
	return s
}
