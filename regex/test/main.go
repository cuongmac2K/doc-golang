package main

import (
	"fmt"
	"regexp"
)

func main() {
	logEntry := "Jan 25 16:58:22 cuongmac2k@gmail.com cqshinns-MacBook-Pro syslogd[115]: ASL Sender Statistics 10.3.251.144 Ha-Noi"

	regex := regexp.MustCompile(`^(?:\S+\s+){10}(\S+)\s+((?:\S+\s+){8}\S+)`)

	match := regex.FindStringSubmatch(logEntry)
	if len(match) > 0 {
		word11 := match[1]
		word12to20 := match[2]
		fmt.Println("Từ thứ 11:", word11)
		fmt.Println("Từ thứ 12 đến 20:", word12to20)
	} else {
		fmt.Println("Không tìm thấy thông tin trong chuỗi log.")
	}
}
