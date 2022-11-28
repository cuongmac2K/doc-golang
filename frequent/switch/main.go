package main

import "fmt"

func main() {
	a := "a"
	switch a {
	case "a":
		fmt.Println("ffff")
		fallthrough
	case "gif":
		fallthrough
	case "png":
		fallthrough
	case "webp":
		fmt.Println("true")
	default:
		fmt.Println("false")
	}
	//fmt.Println("false")
}
