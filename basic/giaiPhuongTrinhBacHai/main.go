package main

import (
	"fmt"
	"math"
)

func main() {
	a := 0
	fmt.Println("nhap a")
	fmt.Scan(&a)
	b := 0
	fmt.Println("nhap b")
	fmt.Scan(&b)
	c := 0
	fmt.Println("nhap c")
	fmt.Scan(&c)
	denta := b*b - 4*a*c
	if denta < 0 {
		fmt.Println("phuong trinh vo nghiem")
	} else if denta == 0 {
		fmt.Println("x1 = x2 = ", -b/(2*a))
	} else {
		x1 := -(float64(b) + math.Sqrt(float64(b))) / (2 * float64(a))
		x2 := -(float64(b) - math.Sqrt(float64(b))) / (2 * float64(a))
		fmt.Println("x1 = ", x1)
		fmt.Println("x2 = ", x2)
	}

}
