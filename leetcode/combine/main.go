package main

import "fmt"

func main() {
	n := 4
	k := 2
	fmt.Println(combine(n, k))
}
func combine(n int, k int) [][]int {
	res := [][]int{}
	var backtrack func(i int, combination []int)
	backtrack = func(start int, combination []int) {
		if len(combination) == k {
			temp := make([]int, k)
			copy(temp, combination)
			res = append(res, temp)
		}
		for i := start; i <= n; i++ {
			backtrack(i+1, append(combination, i))
		}
	}
	backtrack(1, []int{})
	return res
}
