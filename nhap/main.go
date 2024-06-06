package main

import "fmt"

func main() {
	n := 5
	var back func(i int, combination []int)
	back = func(start int, combination []int) {
		fmt.Println("start: ", start)
		fmt.Println(combination)
		for i := start; i <= n; i++ {
			back(i+1, append(combination, i))
		}
	}
	back(1, []int{})
}
