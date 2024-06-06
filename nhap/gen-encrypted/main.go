package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

// Khóa bí mật dùng để ký JWT
var jwtKey = []byte("my_secret_key1111")

// Tạo JWT từ các trường category, name, email
func createJWT(category, name, email string) (string, error) {
	// Tạo claims chứa các thông tin bạn muốn mã hóa
	claims := jwt.MapClaims{
		"category": category,
		"name":     name,
		"email":    email,
	}

	// Tạo token từ các claims và key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ký token và chuyển đổi thành chuỗi
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Giải mã JWT và trả về các trường category, name, email
func decodeJWT(tokenString string) (string, string, string, error) {
	// Giải mã token với key và xác thực phương pháp HS256
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", "", "", err
	}

	// Kiểm tra nếu token hợp lệ
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		category := claims["category"].(string)
		name := claims["name"].(string)
		email := claims["email"].(string)
		return category, name, email, nil
	}

	return "", "", "", fmt.Errorf("invalid token")
}

func main() {
	// Tạo một JWT từ các trường category, name, email
	tokenString, err := createJWT("admin", "John Doe", "johndoe@example.com")
	if err != nil {
		fmt.Println("Error creating JWT:", err)
		return
	}
	fmt.Println("JWT:", tokenString)

	// Kiểm tra lỗi parse JWT
	if err != nil {
		fmt.Println("Error parsing JWT:", err)
		return
	}
	// Giải mã JWT
	category, name, email, err := decodeJWT(tokenString)
	if err != nil {
		fmt.Println("Error decoding JWT:", err)
		return
	}

	fmt.Printf("Decoded: Category=%s, Name=%s, Email=%s\n", category, name, email)
	// Giải mã JWT
	category, name, email, err = decodeJWT1("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjYXRlZ29yeSI6IlNlcnZpY2UgQSIsImVtYWlsIjoiY3VvbmdtYWMya0BtYWlsLmNvbSIsIm5hbWUiOiJhZ2VudCA1NTU1NTU1NTU1NTU1NTU1NSJ9.Q0gALavYvN-hHpiXs4Cuhg0mvCeW2v4a5wJgbu60A7c")
	if err != nil {
		fmt.Println("Error decoding JWT:", err)
		return
	}

	fmt.Printf("Decoded: Category=%s, Name=%s, Email=%s\n", category, name, email)
}

// Giải mã JWT và trả về các trường category, name, email
func decodeJWT1(tokenString string) (string, string, string, error) {
	// Giải mã token với key và xác thực phương pháp HS256
	jwtKey = []byte("9a6432e0-8ddf-46ad-9fcc-b88ec888d3db")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", "", "", err
	}

	// Kiểm tra nếu token hợp lệ
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		category := claims["category"].(string)
		name := claims["name"].(string)
		email := claims["email"].(string)
		return category, name, email, nil
	}

	return "", "", "", fmt.Errorf("invalid token")
}
