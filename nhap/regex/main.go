package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Định nghĩa regex pattern
	//pattern := `[a-f0-9]{8}-?[a-f0-9]{4}-?4[a-f0-9]{3}-?[89ab][a-f0-9]{3}-?[a-f0-9]{12}`
	pattern := `(?!\/.)[A-Za-z0-9_.@]{1,255}$`

	// Chuỗi mẫu để kiểm tra
	//text := "Example UUID: 2a0ce5b9-b05b-4741-a547-68eaa8d738dd and another UUID: b0120bb8-ae50-4cbd-aab7-65835ec0fbd2 b0120bb8-ae50-4cbd-aab7-65835ec0fbd2123e4567-e89b-12d3-a456-426614174000."
	text := "john.doe@domain.com"
	// Compile regex
	re := regexp.MustCompile(pattern)

	// Thay thế tất cả các chuỗi khớp với regex bằng dấu *
	result := re.ReplaceAllString(text, "*")

	// In kết quả
	fmt.Println(result)
}
