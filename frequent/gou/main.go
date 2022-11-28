package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("a ")
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//helo(i)
			fmt.Println(" helo ", i)
		}(i)
	}
	wg.Wait()
}
func helo(i int) {
	fmt.Println(" helo ", i)
}
