package main

import "fmt"

func main() {
	//a := [...]string{"cuong", "la", "ai"}
	//fmt.Println("result la ", a[1])
	//b()
	a()
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}
func c() (i int) {
	defer func() { i++ }()
	return 1
}
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
