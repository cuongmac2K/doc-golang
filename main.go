package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	a := []string{"1", "2"}
	if len(a) > 0 {
		fmt.Println("len a lon hon 0")
	}
}

func g1(a chan string) {
	for i := 1; i <= 10; i++ {
		fmt.Println("1 ", i)
	}
	a <- " 1 cai keo"
	a <- "2 cai keo"
	wg.Done()
}

func g2(a chan string) {
	for i := 11; i <= 20; i++ {
		fmt.Println("2", i)
	}
	fmt.Println("ben 2", <-a)
	wg.Done()
}
