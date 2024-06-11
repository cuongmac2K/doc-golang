package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := fmt.Sprintf("%s/admin/resources", "https://staging.bizflycloud.vn/api/iam/v2")

	body, err := CallAPI("GET", url, nil)
	if err != nil {
		fmt.Println("err 1", err)
	}
	res := map[string]interface{}{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("err 2 ", err)
	}
	data := res["data"]
	res1 := []dataResource{}
	byte, _ := json.Marshal(data)
	err = json.Unmarshal(byte, &res1)
	if err != nil {
		fmt.Println("err 2 ", err)
	}
	mapPatterns := map[string]bool{}
	for _, v := range res1 {
		if len(v.Patterns) > 0 {
			for _, i := range v.Patterns {
				mapPatterns[i] = true
			}
		}
	}
	fmt.Println("\n \n")
	for k, _ := range mapPatterns {
		fmt.Println(k)
	}
}
func CallAPI(method, url string, bodyJSON []byte) ([]byte, error) {
	// Tạo một client HTTP
	client := &http.Client{}

	// Tạo request đến API
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-IAM-TOKEN", "7WNTmylN1EpWLgVwwI4JPi6GqGbZyQifOMoruXY9p7WkpmyuNlmzxlaYHMQv")

	// Gửi request và nhận response
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Đọc nội dung response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type dataResource struct {
	Patterns    []string `json:"patterns"`
	Type        string   `json:"type"`
	Endpoint    string   `json:"endpoint"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
}
