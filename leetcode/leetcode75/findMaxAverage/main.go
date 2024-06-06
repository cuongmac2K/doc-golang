package main

import (
	"fmt"
	"math"
)

func main() {
	//arr := []int{0, 1, 1, 3, 3}
	arr := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage1(arr, k))
}

func findMaxAverage(nums []int, k int) float64 {
	max := math.MinInt
	n := len(nums)
	if n == 1 {
		return float64(nums[0])
	}
	for i := 0; i <= n-k; i++ {
		arr := nums[i : k+i]
		sum := 0
		for _, v := range arr {
			sum += v
		}
		if max < sum {
			max = sum
		}
	}

	return float64(max) / float64(k)
}

func findMaxAverage1(nums []int, k int) float64 {
	ans := math.MinInt
	runningSum := 0
	for idx, val := range nums {
		runningSum += val
		if idx >= k-1 {
			if idx != k-1 {
				runningSum = runningSum - nums[idx-k]
			}
			if ans < runningSum {
				ans = runningSum
			}
		}
		fmt.Println(runningSum)
	}
	return float64(ans) / float64(k)
}
