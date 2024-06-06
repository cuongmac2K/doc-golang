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
	result := [][]int{}
	n := len(nums)
	if n < 4 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < n; k++ {
				for h := k + 1; h < n; h++ {
					if check(nums[i], nums[j], nums[k], nums[h], target) {
						b := []int{nums[i], nums[j], nums[k], nums[h]}
						if checkCoincide(b, result) {
							result = append(result, b)
						}
					}
				}
			}
		}
	}
	return result
}

func check(a, b, c, d, taget int) bool {
	if a+b+c+d == taget {
		return true
	}
	return false
}
func checkCoincide(arr []int, arrArr [][]int) bool {
	//count := 0
	sort.Ints(arr)
	for _, v := range arrArr {
		if v[0] == arr[0] && v[1] == arr[1] && v[2] == arr[2] && v[3] == arr[3] {
			return false
		}

	}
	return true
}
