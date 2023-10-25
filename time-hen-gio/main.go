package main

import (
	"fmt"
	"net/http"
	"time"
)

func task() {
	fmt.Println("Task is running")
}

func main() {
	ticker := time.Tick(1 * time.Minute / 6) // Thiết lập hàm chạy sau mỗi 6 giây

	go func() {
		for range ticker {
			task() // Gọi hàm nhiệm vụ của bạn ở đây
		}
	}()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("chau len ba"))
	})
	http.ListenAndServe(":8000", nil)
}
