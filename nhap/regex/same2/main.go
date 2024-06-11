package main

import (
	"fmt"
	"regexp"
)

func isValidString(s string) bool {
	// Kiểm tra điều kiện đầu tiên: không bắt đầu bằng "/."
	if len(s) >= 2 && s[:2] == "/." {
		return false
	}

	// Kiểm tra điều kiện thứ hai: chỉ chứa các ký tự hợp lệ và độ dài từ 1 đến 255
	re := regexp.MustCompile(`^[A-Za-z0-9_.@]{1,255}$`)
	return re.MatchString(s)
}

func main() {
	testStrings := []string{
		"validString_1@example.com",
		"/.invalidStart",
		"another.valid_string@domain.com",
		"invalid/string",
		"valid_string@example.com",
		"tooLongString....................................................................................................................................................................................@example.com",
	}

	for _, str := range testStrings {
		if isValidString(str) {
			fmt.Printf("Matched: %s\n", str)
		} else {
			fmt.Printf("Did not match: %s\n", str)
		}
	}
}
