package main

import "fmt"

func main() {
	fmt.Println(string(101))
	a := "an_"
	for _, i := range a {
		fmt.Println(i)
	}
	fmt.Println(1 & 1)
}
