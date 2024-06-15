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
	/*	category, name, email, err = decodeJWT1("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjYXRlZ29yeSI6IlNlcnZpY2UgQSIsImVtYWlsIjoiY3VvbmdtYWMya0BtYWlsLmNvbSIsIm5hbWUiOiJhZ2VudCA1NTU1NTU1NTU1NTU1NTU1NSJ9.Q0gALavYvN-hHpiXs4Cuhg0mvCeW2v4a5wJgbu60A7c")
		if err != nil {
			fmt.Println("Error decoding JWT:", err)
			return
		}*/
	fmt.Printf("Decoded: Category=%s, Name=%s, Email=%s\n", category, name, email)
	//decodeJWT2("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjYXRlZ29yeSI6Ik5ldyIsIm5hbWUiOiJzZGsgNTY2NjY2IiwicHJvamVjdF9pZCI6ImIwMTIwYmI4LWFlNTAtNGNiZC1hYWI3LTY1ODM1ZWMwZmJkMiIsInRva2VuX2luX3dvcmtlciI6IjlhNjQzMmUwLThkZGYtNDZhZC05ZmNjLWI4OGVjODg4ZDNkYiIsInR5cGUiOiJzZGsiLCJ1c2VyIjoiY3VvbmdtYWMya0BtYWlsLmNvbSJ9.2cObFwn_i3W-dC1yAoHCkxZNboJPexh47YeVtVrbDsQ")
	//decodeJWT1("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjYXRlZ29yeSI6Ik5ldyIsIm5hbWUiOiJzZGsgNTY2NjY2IiwicHJvamVjdF9pZCI6ImIwMTIwYmI4LWFlNTAtNGNiZC1hYWI3LTY1ODM1ZWMwZmJkMiIsInRva2VuX2luX3dvcmtlciI6IiIsInR5cGUiOiJzZGsiLCJ1c2VyIjoiY3VvbmdtYWMya0BtYWlsLmNvbSJ9.LRfRILawil2sSJwiQ7ELBKNLXsrGFokyLM8kTMiinsM")
	//decodeJWT2("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjYXRlZ29yeSI6Ik5ldyIsIm5hbWUiOiJzZGsgNTY2NjY2IiwicHJvamVjdF9pZCI6ImIwMTIwYmI4LWFlNTAtNGNiZC1hYWI3LTY1ODM1ZWMwZmJkMiIsInRva2VuX2luX3dvcmtlciI6IjlhNjQzMmUwLThkZGYtNDZhZC05ZmNjLWI4OGVjODg4ZDNkYiIsInR5cGUiOiJzZGsiLCJ1c2VyIjoiY3VvbmdtYWMya0BtYWlsLmNvbSJ9.2cObFwn_i3W-dC1yAoHCkxZNboJPexh47YeVtVrbDsQ")
	decodeJWT1("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjYXRlZ29yeSI6Ik5ldyIsIm5hbWUiOiJzZGsgNTY2NjY2IiwicHJvamVjdF9pZCI6ImIwMTIwYmI4LWFlNTAtNGNiZC1hYWI3LTY1ODM1ZWMwZmJkMiIsInRva2VuX2luX3dvcmtlciI6IiIsInR5cGUiOiJzZGsiLCJ1c2VyIjoiY3VvbmdtYWMya0BtYWlsLmNvbSJ9.LRfRILawil2sSJwiQ7ELBKNLXsrGFokyLM8kTMiinsM")

}

// Giải mã JWT và trả về các trường category, name, email
func decodeJWT1(tokenString string) (string, string, string, error) {
	// Giải mã token với key và xác thực phương pháp HS256
	jwtKey = []byte("")
	dataDecodeJwt := DataDecodeJwtSdk{}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", "", "", err
	}

	// Kiểm tra nếu Token hợp lệ
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		category, okC := claims["category"].(string)
		tokenInWorker, okT := claims["token_in_worker"].(string)
		name, okN := claims["name"].(string)
		typeService, okType := claims["type"].(string)
		user, okU := claims["user"].(string)
		projectUuid, okPU := claims["project_id"].(string)
		if okC {
			dataDecodeJwt.Category = category
		}
		if okT {
			dataDecodeJwt.TokenInWorker = tokenInWorker
		}
		if okN {
			dataDecodeJwt.Name = name
		}
		if okType {
			dataDecodeJwt.Type = typeService
		}
		if okU {
			dataDecodeJwt.User = user
		}
		if okPU {
			dataDecodeJwt.ProjectId = projectUuid
		}
		return "", "", "", nil
	}
	return "", "", "", fmt.Errorf("invalid token")
}

type DataDecodeJwtSdk struct {
	Category      string `json:"category"`
	Name          string `json:"name"`
	ProjectId     string `json:"project_id"`
	User          string `json:"user"`
	Type          string `json:"type"`
	TokenInWorker string `json:"token_in_worker"`
}

// Giải mã JWT và trả về các trường category, name, email

func decodeJWT2(tokenString string) (string, string, string, error) {
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
