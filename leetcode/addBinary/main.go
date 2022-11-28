package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "101"
	b := "1111"
	fmt.Println(" ", addBinary(a, b))
}

func addBinary(a string, b string) string {
	num1 := convertStringToArrayInt(a)
	num2 := convertStringToArrayInt(b)
	n1 := len(num1)
	n2 := len(num2)
	fmt.Println(n1, n2)
	result := []int{}
	sodu := 0
	if n1 > 0 && n1 > 0 {

		if n1 < n2 {
			for i := n1 - 1; i >= 0; i-- {
				fmt.Println("===")
				sum, flag := sumBit(num1[i], num2[i], sodu)
				fmt.Println(sum, " and ", flag)
				sodu = 0
				if flag == false {
					result = append(result, sum)
				} else {
					sodu = 1
					result = append(result, sum)
				}
			}
			if sodu == 1 {
				for i := n2 - 1; i >= n1; i-- {
					sum, flag := sumBit(0, num2[i], sodu)
					sodu = 0
					if flag == false {
						result = append(result, sum)
					} else {
						sodu = 1
						result = append(result, sum)
					}
				}
			}
		} else {
			for i := n2 - 1; i >= 0; i-- {
				sum, flag := sumBit(num1[i], num2[i], sodu)
				if flag == false {
					result = append(result, sum)
				} else {
					sodu = 1
					result = append(result, sum)
				}
			}
		}
	} else {
		if len(num1) == 0 {
			return b
		}
		if len(num2) == 0 {
			return a
		}
	}
	fmt.Println(result)
	return ""
}
func convertStringToArrayInt(s string) []int {
	a := []int{}
	char := []rune(s)
	for _, i := range char {
		s1 := string(i)
		i1, _ := strconv.Atoi(s1)
		a = append(a, int(i1))
	}
	return a
}
func sumBit(num1 int, num2 int, sodu int) (sum int, flag bool) {
	if num1+num2+sodu == 0 {
		return 0, false
	}
	if num1+num2+sodu == 1 {
		return 1, false
	}
	if num1+num2+sodu == 2 {
		return 1, true
	}
	if num1+num2+sodu == 3 {
		return 1, true
	}
	fmt.Println("++", num1, num2, sodu)
	return -1, false
}
