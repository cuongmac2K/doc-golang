package main

import "fmt"

func main() {
	height := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height))
}
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	waters := 0
	localWaters := 0
	start := 0

	for i := 1; i < len(height); i++ {
		if height[start] <= height[i] {
			waters += localWaters
			localWaters = 0
			start = i
		} else {
			if i == len(height)-1 {
				localWaters = 0
				end := i
				for j := i - 1; j >= start; j-- {
					if height[end] <= height[j] {
						waters += localWaters
						localWaters = 0
						end = j
					} else {
						localWaters += height[end] - height[j]
					}
				}
			} else {
				localWaters += height[start] - height[i]
			}
		}
	}

	return waters
}
