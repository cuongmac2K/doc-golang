package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Giả sử đây là mật khẩu đã băm được lưu trữ trong cơ sở dữ liệu
	hashedPassword := "$2a$10$IRn4n16aSCoC5jmDRvPLGuTVKt8ZXijRTYZrp7g9SMAEaow3f7cau"

	// Mật khẩu người dùng nhập vào
	inputPassword := "your_input_password"

	// So sánh mật khẩu nhập vào với mật khẩu đã băm
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	if err != nil {
		fmt.Println("Password does not match")
	} else {
		fmt.Println("Password matches")
	}
}
