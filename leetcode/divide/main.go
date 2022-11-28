package main

import "fmt"

func main() {
	dividend := 10
	divisor := 0
	fmt.Println(divide(dividend, divisor))
}
func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return 0
	}
	return dividend / divisor
}
