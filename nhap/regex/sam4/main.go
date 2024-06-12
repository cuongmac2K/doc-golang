package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	url := "https://hn-staging2.bizflycloud.vn/api/iaas-cloud/keypairs/string_id_123ggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"
	data, tg := CheckPatternKeyPairs(url)
	if tg {
		fmt.Println("Transformed URL:", data)
	} else {
		fmt.Println("false")
	}
}

func CheckPatternKeyPairs(url string) (string, bool) {
	if !strings.Contains(url, "/keypairs/") {
		return "", false
	}
	// Định nghĩa regex pattern
	pattern := `^[A-Za-z0-9_.@]{1,255}$`
	re := regexp.MustCompile(pattern)

	// Tìm vị trí cuối cùng của dấu "/"
	lastSlashIndex := strings.LastIndex(url, "/")

	if lastSlashIndex != -1 {
		// Tách phần cuối cùng của URL
		lastPart := url[lastSlashIndex+1:]

		// Kiểm tra nếu phần cuối không bắt đầu với "/"
		if !strings.HasPrefix(lastPart, "/") && re.MatchString(lastPart) {
			// Thay thế phần cuối bằng "*"
			url = url[:lastSlashIndex+1] + "*"
		}
	}
	return url, true
}
