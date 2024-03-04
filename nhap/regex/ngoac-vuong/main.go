package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "2024/02/19 09:39:08 [cuongmac2k@gmail.com]-[Hà Đông - Hà Nội]-[192.168.0.1]"

	// Loại bỏ các ký tự ngoặc nhọn từ chuỗi
	data = strings.Trim(data, "[]")
	time := data[:20]
	data = data[20:]
	fmt.Printf("Thời gian: %s\n", time)
	// Chia chuỗi thành các phần tử tại dấu "-]"
	elements := strings.Split(data, "]-[")

	// In phần tử theo định dạng
	fmt.Printf("Email: %s\n", elements[0])
	fmt.Printf("Địa chỉ: %s\n", elements[1])
	fmt.Printf("IP: %s\n", elements[2])
}
