package main

import (
	"fmt"
)

func main() {
	num := 13
	fmt.Println(intToRoman(num))
}

func intToRoman(num int) string {
	var r string
	v := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	s := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < 13; i++ {
		for num >= v[i] {
			r += s[i]
			num -= v[i]
		}
	}
	return r
}
