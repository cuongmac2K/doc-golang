package main

// 1
import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

// 2
func hello(name string, i int) {
	message := fmt.Sprintf("Hi, %v %v ", name, i)
	fmt.Println(message)
}

func runCronJobs() {
	// 3
	s := gocron.NewScheduler(time.UTC)
	i := 0
	// 4
	s.Every(100).Milliseconds().Do(func() {
		hello("John Doe", i)
		i++
	})

	// 5
	s.StartBlocking()
}

// 6
func main() {
	runCronJobs()
}
