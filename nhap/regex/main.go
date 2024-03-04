package main

import (
	"fmt"
	"regexp"
)

func main() {
	logEntry := "2024/02/01 15:00:08 abc@gmail.com cqshinns-MacBook-Pro syslogd[115]: ASL Sender Statistics 10.3.251.144 Ha-Noi19"

	regex := regexp.MustCompile(`(\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}) ([A-Za-z0-9._%+-]+@gmail\.com) ([A-Za-z0-9-]+) (.*) (\d+\.\d+\.\d+\.\d+) (.*)`)

	match := regex.FindStringSubmatch(logEntry)
	if len(match) > 0 {
		dateTime := match[1]
		email := match[2]
		name := match[3]
		message := match[4]
		ip := match[5]
		address := match[6]
		fmt.Println("Ngày tháng và thời gian:", dateTime)
		fmt.Println("Email:", email)
		fmt.Println("Tên máy:", name)
		fmt.Println("Chuỗi tin nhắn:", message)
		fmt.Println("Địa chỉ IP:", ip)
		fmt.Println("Địa chỉ:", address)
	} else {
		fmt.Println("Không tìm thấy thông tin trong chuỗi log.")
	}
}
