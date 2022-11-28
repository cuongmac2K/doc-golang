package main

import "fmt"

func defFoo() {
	fmt.Println("defFoo() started")

	if r := recover(); r != nil {
		fmt.Println("WOHA! Program is panicking with value", r)
	}

	fmt.Println("defFoo() done=========")
}

func normMain() {

	fmt.Println("normMain() started")

	defer defFoo() // defer defFoo call

	panic("oke len lele") // panic here

	fmt.Println("normMain() done")
}

func main() {
	fmt.Println("main() started")

	normMain() // normal call

	fmt.Println("main() done")
}
