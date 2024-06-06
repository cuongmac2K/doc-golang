package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "the sky is blue"
	s = "  hello   world  "
	s = "  t"
	fmt.Println(reverseWords(s))
}
func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	arr := strings.Fields(s)
	n := len(arr)
	if n == 0 || n == 1 {
		return s
	}
	// reverse
	a := []string{}
	for i := n - 1; i >= 0; i-- {
		a = append(a, arr[i])
	}
	return strings.Join(a, " ")
}

func reverseWords1(s string) string {
	// Bước 1: Loại bỏ khoảng trắng ở đầu và cuối
	start := 0
	end := len(s) - 1

	// Loại bỏ khoảng trắng ở đầu
	for start <= end && unicode.IsSpace(rune(s[start])) {
		start++
	}

	// Loại bỏ khoảng trắng ở cuối
	for end >= start && unicode.IsSpace(rune(s[end])) {
		end--
	}

	// Nếu chuỗi rỗng sau khi loại bỏ khoảng trắng, trả về chuỗi rỗng
	if start > end {
		return ""
	}

	// Bước 2: Đảo ngược thứ tự các từ trong chuỗi
	var result []byte
	wordEnd := end + 1

	for i := end; i >= start; i-- {
		if unicode.IsSpace(rune(s[i])) {
			if wordEnd > i+1 {
				result = append(result, s[i+1:wordEnd]...)
				result = append(result, ' ')
			}
			wordEnd = i
		}
	}

	// Thêm từ cuối cùng (hoặc từ duy nhất nếu không có khoảng trắng)
	result = append(result, s[start:wordEnd]...)

	return string(result)
}
