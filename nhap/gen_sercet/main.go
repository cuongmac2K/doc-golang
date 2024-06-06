package main

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	a := uuid.New().String()
	fmt.Println(a)
	encoded := base64.StdEncoding.EncodeToString([]byte(a))
	fmt.Println("Encoded:", encoded)
}
