package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

// DataDecodeJwtSdk represents the structure of the JWT payload
type DataDecodeJwtSdk struct {
	Category      string
	Name          string
	Type          string
	User          string
	ProjectId     string
	TokenInWorker string
}

var jwtKey = []byte("9a6432e0-8ddf-46ad-9fcc-b88ec888d3db")

func CreateJWTSdk(decodeJwt DataDecodeJwtSdk) (string, error) {
	// Tạo claims chứa các thông tin bạn muốn mã hóa
	claims := jwt.MapClaims{
		"category":        decodeJwt.Category,
		"name":            decodeJwt.Name,
		"type":            decodeJwt.Type,
		"user":            decodeJwt.User,
		"project_id":      decodeJwt.ProjectId,
		"token_in_worker": decodeJwt.TokenInWorker,
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

func decodeJWT1(tokenString string) (DataDecodeJwtSdk, error) {
	// Giải mã token với key và xác thực phương pháp HS256
	var dataDecodeJwt DataDecodeJwtSdk
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return dataDecodeJwt, err
	}

	// Kiểm tra nếu Token hợp lệ
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if category, ok := claims["category"].(string); ok {
			dataDecodeJwt.Category = category
		}
		if tokenInWorker, ok := claims["token_in_worker"].(string); ok {
			dataDecodeJwt.TokenInWorker = tokenInWorker
		}
		if name, ok := claims["name"].(string); ok {
			dataDecodeJwt.Name = name
		}
		if typeService, ok := claims["type"].(string); ok {
			dataDecodeJwt.Type = typeService
		}
		if user, ok := claims["user"].(string); ok {
			dataDecodeJwt.User = user
		}
		if projectUuid, ok := claims["project_id"].(string); ok {
			dataDecodeJwt.ProjectId = projectUuid
		}
		return dataDecodeJwt, nil
	}
	return dataDecodeJwt, fmt.Errorf("invalid token")
}

func main() {
	// Example usage
	jwtData := DataDecodeJwtSdk{
		Category:      "New",
		Name:          "sdk 566666",
		Type:          "sdk",
		User:          "cuongmac2k@mail.com",
		ProjectId:     "b0120bb8-ae50-4cbd-aab7-65835ec0fbd2",
		TokenInWorker: string(jwtKey),
	}

	token, err := CreateJWTSdk(jwtData)
	if err != nil {
		fmt.Println("Error creating JWT:", err)
		return
	}
	fmt.Println("Generated Token:", token)

	//decodedData, err := decodeJWT1(token)
	decodedData, err := decodeJWT1("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjYXRlZ29yeSI6Ik5ldyIsIm5hbWUiOiJzZGsgNTY2NjY2IiwicHJvamVjdF9pZCI6ImIwMTIwYmI4LWFlNTAtNGNiZC1hYWI3LTY1ODM1ZWMwZmJkMiIsInRva2VuX2luX3dvcmtlciI6IjlhNjQzMmUwLThkZGYtNDZhZC05ZmNjLWI4OGVjODg4ZDNkYiIsInR5cGUiOiJzZGsiLCJ1c2VyIjoiY3VvbmdtYWMya0BtYWlsLmNvbSJ9.2cObFwn_i3W-dC1yAoHCkxZNboJPexh47YeVtVrbDsQ")
	if err != nil {
		fmt.Println("Error decoding JWT:", err)
		return
	}
	fmt.Println("Decoded Data:", decodedData)
}
