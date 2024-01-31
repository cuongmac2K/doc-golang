package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	for i := 0; i < 10; i++ {
		// Gọi hàm để thực hiện API POST
		err := callAPI("http://localhost:8080/logs", map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		}, map[string]string{
			"Content-Type": "application/json",
			"X-cuong":      "abc-Id_1000",
		})

		if err != nil {
			fmt.Println("Lỗi khi gọi API:", err)
		}
		//time.Sleep(time.Second / 100)
	}
}

func callAPI(url string, bodyData map[string]interface{}, headers map[string]string) error {
	// Chuyển đổi bodyData thành JSON
	bodyJSON, err := json.Marshal(bodyData)
	if err != nil {
		return err
	}

	// Tạo request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return err
	}

	// Thiết lập các header
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Gửi request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Kiểm tra mã trạng thái của response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Mã trạng thái lỗi: %d", resp.StatusCode)
	}

	// Xử lý response...
	fmt.Println("Gọi API thành công")

	return nil
}
