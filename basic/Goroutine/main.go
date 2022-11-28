package main

import (
	"fmt"
	"sync"
	"time"
)

func go1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		fmt.Println("i la ", i)
	}
	fmt.Print("\nI am first runner")

}
func go2(wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < 500; j++ {
		fmt.Println("a la ", j)
	}
	fmt.Print("\nI am second runner")
}
func go3(wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < 500; j++ {
		fmt.Println("a la ", j)
	}
	fmt.Print("\nI am second runner")
}
func go4(wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < 500; j++ {
		fmt.Println("a la ", j)
	}
	fmt.Print("\nI am second runner")
}
func main() {
	start := time.Now()
	wg := new(sync.WaitGroup)
	wg.Add(4)
	go go1(wg)
	go go2(wg)
	go go3(wg)
	go go4(wg)
	wg.Wait()

	fmt.Println(" thoi gian chay ", time.Since(start))

}
