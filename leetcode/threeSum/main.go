package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
func threeSum(nums []int) [][]int {
	a := [][]int{}
	n := len(nums)
	if n < 3 {
		return a
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if BaSo(nums[i], nums[j], nums[k]) && checkExit(nums[i], nums[j], nums[k], a) {
					a = append(a, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return a
}

func BaSo(a, b, c int) bool {
	if a+b+c == 0 {
		return true
	}
	return false
}
func checkExit(a, b, c int, arr [][]int) bool {
	for i := 0; i < len(arr); i++ {
		sort.Ints(arr[i])
		k := []int{a, b, c}
		sort.Ints(k)
		if k[0] == arr[i][0] && k[1] == arr[i][1] && k[2] == arr[i][2] {
			return false
		}
	}
	return true
}
