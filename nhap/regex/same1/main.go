package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	//baseURL := "https://hn-staging2.bizflycloud.vn/api/_/abcd1234efgh5678/0123456789ABCDEFabcdef012345/123e4567-e89b-42d3-a456-426614174000"
	//
	//// Các mẫu biểu thức chính quy
	//patterns := []string{
	//	"[a-z0-9]{16}",
	//	"[0-9a-fA-F]{24}",
	//	"[a-f0-9]{8}-?[a-f0-9]{4}-?4[a-f0-9]{3}-?[89ab][a-f0-9]{3}-?[a-f0-9]{12}",
	//}

	baseURL := "https://hn-staging2.bizflycloud.vn/api/_/cs/123e4567-e89b-42d3-a456-426614174000/123e4567-e89b-42d3-a456-426614174000"

	// Các mẫu biểu thức chính quy
	patterns := []string{
		//"[a-z0-9]{16}",
		"[a-f0-9]{8}-?[a-f0-9]{4}-?4[a-f0-9]{3}-?[89ab][a-f0-9]{3}-?[a-f0-9]{12}",
		"[a-f0-9]{8}-?[a-f0-9]{4}-?4[a-f0-9]{3}-?[89ab][a-f0-9]{3}-?[a-f0-9]{12}",
	}

	// Biên dịch các biểu thức chính quy
	var regexps []*regexp.Regexp
	for _, pattern := range patterns {
		regexps = append(regexps, regexp.MustCompile(pattern))
	}

	// Tạo slice từ chuỗi ban đầu bằng cách tách các phần theo dấu /
	parts := strings.Split(baseURL, "/")

	// Duyệt qua các phần và thực hiện thay thế tương ứng
	for i, part := range parts {
		for _, re := range regexps {
			if re.MatchString(part) {
				parts[i] = "*"
				break
			}
		}
	}

	// Ghép lại các phần thành chuỗi mới
	newURL := strings.Join(parts, "/")
	fmt.Println("New URL:", newURL)
}
