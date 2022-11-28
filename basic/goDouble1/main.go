package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Done()
	go go1(mu)
	wg.Wait()
	go go2(mu)
	wg.Wait()
	fmt.Println("thoi gian chay la ", time.Since(start))

}
func go1(mu sync.Mutex) {
	mu.Lock()
	for i := 0; i < 500; i++ {
		fmt.Println("i la ", i)
	}
	mu.Unlock()
}
func go2(mu sync.Mutex) {
	mu.Lock()
	for i := 0; i < 500; i++ {
		fmt.Println("a la ", i)
	}
	mu.Unlock()
}
