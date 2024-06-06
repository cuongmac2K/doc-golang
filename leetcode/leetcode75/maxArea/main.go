package main

import (
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
func maxArea(height []int) int {
	n := len(height)
	maxAre := 0
	if n == 1 {
		return 0
	}
	if n == 2 {
		if height[0] < height[1] {
			return height[0]
		}
		return height[1]
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			area := 0
			if height[i] <= height[j] {
				area = height[i] * (j - i)
			} else {
				area = height[j] * (j - i)
				fmt.Println(height[i], "and ", (i - i))
			}

			if area > maxAre {
				maxAre = area
			}
		}
	}
	return maxAre
}
