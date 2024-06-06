package main

import (
	"fmt"
	"time"
)

func main() {
	// Chuỗi đại diện cho thời gian
	timeLog := "2024-04-17 12:09:30"

	// Định dạng mẫu của chuỗi thời gian
	layout := "2006-01-02 15:04:05"

	// Phân tích chuỗi thời gian thành kiểu time.Time
	timeObjLog, err := time.Parse(layout, timeLog)
	if err != nil {
		fmt.Println("Lỗi khi phân tích chuỗi thời gian:", err)
		return
	}
	now := time.Now()
	threeHourAgo := now.Add(-11 * time.Hour)
	fmt.Println("Thời gian timeObjLog:", timeObjLog)
	fmt.Println("Thời gian :", threeHourAgo)
	// In ra kiểu thời gian đã được chuyển đổi

	if threeHourAgo.Before(timeObjLog) {
		fmt.Println("push loki")
	}

}
