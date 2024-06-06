package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Dữ liệu chứa các dấu cách và xuống dòng
		data := "\n[INPUT]\n    Name        tail\n    Path        {path_file_log}\n\n[OUTPUT]\n    Name        http\n    Match       *\n    Host        0.0.0.0\n    Port        6060\n    Format      json_lines\n    Header      X-Activity-Token 1713178615:8014aeaa20d8c9679f64910fec0c688f88dbfb5570872484161c059adea51c5d\n    Header      Project-Id 1c9a6a4b-bd7d-4ae7-acf1-691b09f21a76\n    Header      Uuid 19b4f521-ae0e-4e79-ae54-3a4135d4210e\n    URI         /logs/create\n    Workers     1\n"
		// Thiết lập header cho response là plain text
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		// Gửi dữ liệu về client
		_, err := w.Write([]byte(data))
		if err != nil {
			http.Error(w, "Lỗi trong quá trình gửi dữ liệu", http.StatusInternalServerError)
			return
		}
	})

	// Khởi động server trên cổng 8080
	http.ListenAndServe(":8081", nil)
}
