package main

import (
	"fmt"
	"math"
)

func main() {
	x := float64(1.000001)
	n := 2147483647
	fmt.Println(myPow(x, n))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	a := float64(1)
	if n > 0 {
		for i := n; i > 0; i-- {
			a *= x
		}
	} else {
		for i := math.Abs(float64(n)); i > 0; i-- {
			a *= x
		}
		a = float64(1) / a
	}
	return a
}
