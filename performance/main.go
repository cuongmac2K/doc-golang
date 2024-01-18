package main

import (
	"fmt"
	"time"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	start := time.Now()

	func() {
		fmt.Println("Fib(30) =", fib(30))
	}()

	func() {
		fmt.Println("Fib(31) =", fib(31))
	}()

	//time.Sleep(5 * time.Second) // Đợi 5 giây để quan sát kết quả
	elapsed := time.Since(start)
	fmt.Printf("Concurrency execution time: %s\n", elapsed)
}
