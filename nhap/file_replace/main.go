package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Đường dẫn tệp với đường dẫn tuyệt đối
	filePathAbsolute := "/home/cuong/Desktop/data/code/golang/doc-golang/nhap/file_replace/temple.txt"

	// Đọc nội dung của tệp với đường dẫn tuyệt đối
	contentAbsolute, err := ioutil.ReadFile(filePathAbsolute)
	if err != nil {
		log.Fatalf("Không thể đọc file tuyệt đối: %v", err)
	}

	// Đọc nội dung của tệp với đường dẫn tương đối

	// Chuyển đổi dữ liệu từ byte slice thành chuỗi
	fileContentAbsolute := string(contentAbsolute)

	// replace data in file
	s := strings.Replace(fileContentAbsolute, "Config_host_activity", "0.0.0.0", -1)
	s = strings.Replace(s, "Config_port_activity", "8080", -1)
	s = strings.Replace(s, "Config_Activity_Token", "hhhhhh", -1)
	s = strings.Replace(s, "Config_Project_id", "456", -1)
	s = strings.Replace(s, "Config_Uuid", "123", -1)

	// In ra nội dung của tệp
	fmt.Println("Nội dung của tệp với đường dẫn tuyệt đối:")
	fmt.Println(s)

}
