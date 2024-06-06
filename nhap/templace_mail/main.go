package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

// Struct để đại diện cho dữ liệu trong mẫu email
type FileConfigData struct {
	ConfigHost      string
	ConfigPost      string
	ConfigToken     string
	ConfigProjectId string
	ConfigUUid      string
}

func main() {
	// Tạo dữ liệu cho mẫu email
	data := FileConfigData{
		ConfigHost:      "0.0.0.0",
		ConfigPost:      "8080",
		ConfigToken:     "abc",
		ConfigProjectId: "ddddd",
		ConfigUUid:      "cdcdcd",
	}

	// Định nghĩa mẫu email
	const configTemplate = `
[INPUT]
    Name        tail
    Path        {path_file_log}

[OUTPUT]
    Name        http
    Match       *
    Host        {{.ConfigHost}}
    Port        {{.ConfigPost}}
    Format      json_lines
    Header      X-Activity-Token {{.ConfigToken}}
    Header      Project-Id {{.ConfigProjectId}}
    Header      Uuid {{.ConfigUUid}}
    URI         /logs/create
    Workers     1
`

	// Tạo một template mới từ mẫu văn bản
	tmpl, err := template.New("").Parse(configTemplate)
	if err != nil {
		log.Fatalf("Không thể tạo mẫu email: %v", err)
	}

	// Tạo một buffer để lưu trữ kết quả của mẫu
	var emailContent bytes.Buffer
	err = tmpl.Execute(&emailContent, data)
	if err != nil {
		log.Fatalf("Không thể điền dữ liệu vào mẫu email: %v", err)
	}
	fmt.Println(emailContent.String())
}

// Hàm giả lập gửi email (không gửi thực sự trong ví dụ này)
func sendEmail(emailContent string) {
	fmt.Println("Đã gửi email đi:", emailContent)
	// Các bước gửi email thực tế sẽ được thực hiện ở đây
}
