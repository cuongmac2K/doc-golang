package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func main() {
	mu.Lock()
	go func() {
		for j := 0; j < 10; j++ {
			fmt.Println("a la ", j)
			go1()
		}
	}()
	mu.Unlock()
	time.Sleep(500 * time.Millisecond)
}
func go1() {
	//mu.Lock()
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("i la ", i)
		}
	}()
	//	mu.Unlock()
}
