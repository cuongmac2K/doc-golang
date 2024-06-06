package main

import (
	"fmt"
	"strings"
)

// Hàm để chuẩn hóa chuỗi
func normalizeSpaces(s string) string {
	// Dùng hàm Fields để tách chuỗi thành các từ, tự động loại bỏ các khoảng trắng thừa
	words := strings.Fields(s)
	// Nối các từ lại với nhau bằng một khoảng trắng duy nhất
	return strings.Join(words, " ")
}

func main() {
	// Ví dụ chuỗi cần chuẩn hóa
	s := "  Đây   là   một   ví   dụ   về   chuỗi   có   nhiều   khoảng trắng  "
	normalized := normalizeSpaces(s)
	fmt.Println(normalized) // Output: "Đây là một ví dụ về chuỗi có nhiều khoảng trắng"
}
