package main

import (
	"fmt"
	"sort"
)

func main() {

	num1 := []int{1, 3}
	num2 := []int{2, 4, 5}
	fmt.Println(medianAray(num1, num2))
}
func medianAray(num1 []int, num2 []int) float64 {
	totallen := len(num1) + len(num2)
	num3 := make([]int, totallen)
	var median float64
	// combing 2 arrays
	var c = 0
	for i := 0; i < totallen; i++ {
		if i < len(num1) {
			num3[i] = num1[i]
		} else if i >= len(num1) && c < len(num2) {
			num3[i] = num2[c]
			c++
		}
	}
	sort.Ints(num3)
	if totallen%2 == 0 {
		median = float64(num3[totallen/2-1]+num3[totallen/2]) / 2.0
		return median
	}
	median = float64(num3[totallen/2])
	return median
}
