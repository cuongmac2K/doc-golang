package main

import (
	"fmt"
	"time"
)

func main() {

	// Chuyển định dạng thời gian
	formattedTime := time.Now().UTC().Format("2006-01-02 15:04:05")
	fmt.Println("Thời gian đã chuyển đổi:", formattedTime)
}
