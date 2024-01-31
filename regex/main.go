package main

import (
	"fmt"
	"regexp"
)

var regex_abc_service = "(\\w{3} \\d{2} \\d{2}:\\d{2}:\\d{2}) (\\b[A-Za-z0-9._%+-]+@gmail\\.com\\b) (\\w+-\\w+-\\w+) (.*) (\\d+\\.\\d+\\.\\d+\\.\\d+) (.*)"

func main() {
	logEntry := "Jan 25 16:58:22 cuongmac2k@gmail.com cqshinns-MacBook-Pro syslogd[115]: ASL Sender Statistics 10.3.251.144 Ha-Noi"

	regex := regexp.MustCompile(regex_abc_service)

	match := regex.FindStringSubmatch(logEntry)
	if len(match) > 0 {
		dateTime := match[1]
		email := match[2]
		name := match[3]
		message := match[4]
		ip := match[5]
		address := match[6]
		fmt.Println("Ngày tháng và thời gian:", dateTime)
		fmt.Println("email:", email)
		fmt.Println("Tên máy:", name)
		fmt.Println("Chuỗi tin nhắn:", message)
		fmt.Println("Địa chỉ IP:", ip)
		fmt.Println("Địa chỉ:", address)
	} else {
		fmt.Println("Không tìm thấy thông tin trong chuỗi log.")
	}
}
