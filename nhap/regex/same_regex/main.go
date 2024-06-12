package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	url := "https://hn-staging2.bizflycloud.vn/api/_/abcd1234efgh5678/0123456789ABCDEFabcdef012345/123e4567-e89b-42d3-a456-426614174000"
	patterns := []string{
		"[a-z0-9]{16}",
		"[0-9a-fA-F]{24}",
		"[a-f0-9]{8}-?[a-f0-9]{4}-?4[a-f0-9]{3}-?[89ab][a-f0-9]{3}-?[a-f0-9]{12}",
	}

	//url := "https://hn-staging2.bizflycloud.vn/api/domains/123e4567-e89b-42d3-a456-426614174000/namespaces/ahhhhhhhbc/directories"
	//patterns := []string{
	//	"namespaces/(.*?)/directories",
	//	"[a-f0-9]{8}-?[a-f0-9]{4}-?4[a-f0-9]{3}-?[89ab][a-f0-9]{3}-?[a-f0-9]{12}",
	//}

	//url := "https://hn-staging2.bizflycloud.vn/api/domains/123e4567-e89b-42d3-a456-426614174000/namespaces/ahhhhhhhbc/copy-filesf"
	//patterns := []string{
	//	"namespaces/(.*?)/copy-files",
	//	"[a-f0-9]{8}-?[a-f0-9]{4}-?4[a-f0-9]{3}-?[89ab][a-f0-9]{3}-?[a-f0-9]{12}",
	//}

	// Biên dịch các biểu thức chính quy
	var regexps []*regexp.Regexp
	for _, pattern := range patterns {
		regexps = append(regexps, regexp.MustCompile(pattern))
	}

	originalURL := url

	for _, re := range regexps {
		switch {
		case strings.Contains(re.String(), "(.*?)"):
			// Thay thế toàn bộ phần khớp nếu mẫu là namespaces/(.*?)/directories
			stringsRe := strings.Replace(re.String(), "(.*?)", "*", 1)
			url = re.ReplaceAllString(url, stringsRe)
		default:
			parts := strings.Split(url, "/")
			for i, part := range parts {
				if re.MatchString(part) {
					parts[i] = "*"
					break
				}
			}
			url = strings.Join(parts, "/")
		}
	}

	// In ra URL gốc và URL đã thay thế
	fmt.Printf("Original URL: %s\nReplaced URL: %s\n\n", originalURL, url)

}
