package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//now := time.Now()
	//processNom(100)
	//fmt.Println("thoi gian chay", time.Since(now))

	now := time.Now()
	processChannel(100)
	fmt.Println("thoi gian chay", time.Since(now))
}
func processNom(number int) {
	for i := 0; i < number; i++ {
		fmt.Println(fib(i))
	}
}
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
func processChannel(number int) {
	numberOfWorker := runtime.NumCPU()
	jobs := make(chan int, number)
	results := make(chan int, number)
	for i := 0; i < numberOfWorker; i++ {
		go worker(jobs, results)
	}
	for i := 1; i <= number; i++ {
		jobs <- i
	}
	close(jobs)
	for j := 0; j < number; j++ {
		fmt.Println(<-results)
	}
}
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}
