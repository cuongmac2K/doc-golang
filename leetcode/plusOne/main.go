package main

import "fmt"

func main() {
	a := []int{1, 2, 9}
	fmt.Println(plusOne(a))
}
func plusOne(digits []int) []int {
	if digits[len(digits)-1]+1 == 10 {
		digits[len(digits)-1] = 1
		digits = append(digits, 0)
	} else {
		digits[len(digits)-1] = digits[len(digits)-1] + 1
	}
	return digits
}
