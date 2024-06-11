package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := "http:/domains/abc/namespaces/foo/unzip-files"
	re := regexp.MustCompile("namespaces/(.*?)/unzip-files")

	// Thay thế tất cả các chuỗi khớp với regex bằng dấu *
	result := re.ReplaceAllString(str, "*")
	//fmt.Println("reslut sau khi thay *: ", result)
	fmt.Println(result)
	url := "/domains/*/namespaces/*/unzip-files"
	// convert url to form
	urlForm := re.ReplaceAllString(url, "*")
	fmt.Println(urlForm)
	if strings.Contains(result, "") {

	}
}
