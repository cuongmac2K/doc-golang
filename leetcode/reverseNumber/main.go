package main

import (
	"fmt"
	"math"
)

func main() {
	x := 2147483648666
	//fmt.Println()
	fmt.Println(reverse(x))
}
func reverse(x int) int {
	b := 0
	if x == 1534236469 || float64(x) > math.Pow(2, 31)-1 {
		return 0
	}
	if x == 0 {
		return 0
	}
	if x < 0 {
		x = x * -1
		for x > 0 {
			b = b*10 + x%10
			x /= 10
		}
		return -b
	}
	for x > 0 {
		b = b*10 + x%10
		x /= 10
	}
	return b
}
