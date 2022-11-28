package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int32
	for i := 0; i < 99999999; i++ {
		var wg sync.WaitGroup
		wg.Add(1)
		go func(wg *sync.WaitGroup, count *int32) {
			defer wg.Done()
			atomic.AddInt32(count, 1)
		}(&wg, &count)
		wg.Wait()
	}
	fmt.Println(" count la ", count)

}
