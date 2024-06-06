package main

import "fmt"

func main() {
	candies := []int{2, 3, 5, 1, 3}
	//candies := []int{4, 2, 1, 1, 2}
	n := 3
	fmt.Println(kidsWithCandies(candies, n))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, i := range candies {
		if i > max {
			max = i
		}
	}
	n := len(candies)
	result := []bool{}
	for i := 0; i < n; i++ {
		if candies[i]+extraCandies >= max {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}

//func kidsWithCandies(candies []int, extraCandies int) []bool {
//	biggest := 0
//	for _, candy := range candies {
//		if candy > biggest {
//			biggest = candy
//		}
//	}
//	potential := make([]bool, len(candies))
//	for i, candy := range candies {
//		potential[i] = candy+extraCandies >= biggest
//	}
//	return potential
//}
