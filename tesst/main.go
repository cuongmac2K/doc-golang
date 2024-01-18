package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Getenv("a")
	if a != "chau " ||
		a == "chau len ba " {
		fmt.Println("ok")
	}
}
