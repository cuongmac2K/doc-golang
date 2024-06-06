package main

import (
	"fmt"
	"time"
)

func main() {
	// Lấy thời điểm hiện tại
	now := time.Now()

	// Thời điểm trong quá khứ (5 phút trước)
	pastTime := now.Add(-5 * time.Minute)

	// Kiểm tra nếu `now` trước `pastTime`
	if now.Before(pastTime) {
		fmt.Println("Thời điểm hiện tại là trước thời điểm trong quá khứ (5 phút trước)")
	} else {
		fmt.Println("Thời điểm hiện tại là sau hoặc bằng thời điểm trong quá khứ (5 phút trước)")
	}
}
