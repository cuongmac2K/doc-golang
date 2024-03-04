package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	fileName := "output.txt" // Tên của file để lưu kết quả

	// Tạo hoặc mở file để ghi
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Lỗi khi mở file:", err)
		return
	}
	defer file.Close()
	fmt.Println("abc")
	// Tạo cmd để in ra dòng văn bản
	cmd := exec.Command("echo", "-e", "Dòng văn bản thứ nhất")
	cmd.Stdout = file
	if err := cmd.Run(); err != nil {
		// Xử lý lỗi
		fmt.Println("Lỗi khi chạy lệnh:", err)
	}

	// Đợi một khoảng thời gian
	time.Sleep(1 * time.Second)

	// Tạo cmd khác để in ra dòng văn bản khác
	cmd2 := exec.Command("echo", "-e", "Dòng văn bản thứ hai")
	cmd2.Stdout = file
	if err := cmd2.Run(); err != nil {
		// Xử lý lỗi
		fmt.Println("Lỗi khi chạy lệnh:", err)
	}

	// Hoàn thành ghi file
	file.Sync()
	fmt.Println("Hoàn thành ghi file:", fileName)
}
