package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func gentoken(secret, projectUUID, UUID string) string {
	// Tạo timestamp hiện tại (ví dụ: sử dụng Unix time)
	timestamp := time.Now().Unix()

	// Tạo chuỗi dữ liệu cần ký
	data := fmt.Sprintf("%d:%s:%s", timestamp, projectUUID, UUID)

	// Tạo HMAC từ chuỗi dữ liệu và secret key
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))

	// Kết hợp timestamp và chữ ký để tạo token
	token := fmt.Sprintf("%d:%s", timestamp, signature)

	return token
}

func main() {
	// Các thông tin đầu vào
	secret_gen_token := "951bfbe9-518e-488b-bdfd-eaf80edf9444"
	projectUUID := "964928c6-fed4-4a24-a6ac-2b5daeaab626"
	UUID := "896eeb36-a7bf-46da-8f36-3f86e8c67b47"

	// Gọi hàm gentoken để tạo token
	token := gentoken(secret_gen_token, projectUUID, UUID)

	// In ra token được tạo
	fmt.Println("Generated Token:", token)
	fmt.Println("len Token:", len(token))
}
