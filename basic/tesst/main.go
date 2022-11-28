package main

import "fmt"

func main() {
	requestBody := map[string]interface{}{
		"ids": []string{"abc", "def"},
		"tag": "tag thu 1",
	}
	fmt.Println(requestBody)
}
