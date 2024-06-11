package main

import (
	"fmt"
	"regexp"
)

func isValidString(s string) bool {
	// Kiểm tra điều kiện đầu tiên: không bắt đầu bằng "/"
	if len(s) > 0 && s[0] == '/' {
		return false
	}
	// Kiểm tra điều kiện thứ hai: chỉ chứa các ký tự hợp lệ và độ dài từ 1 đến 255
	re := regexp.MustCompile(`^[A-Za-z0-9_.@]{1,255}$`)
	return re.MatchString(s)
}

func main() {
	fmt.Println(isValidString("example.user"))                              // True
	fmt.Println(isValidString("/invalidStart"))                             // False
	fmt.Println(isValidString("john.doe@domain.com"))                       // True
	fmt.Println(isValidString("A1._@valid"))                                // True
	fmt.Println(isValidString("tooLongString" + string(make([]byte, 240)))) // False, quá dài
}
