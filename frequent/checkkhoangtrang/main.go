package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "{\"minh_oke\":\"string\",\"value\":\"HNffff\"}"
	s := strings.Split(str, ":")
	fmt.Println(s)
	s1 := strings.Split(s[2], "\"")
	fmt.Println(s1[1])
}
