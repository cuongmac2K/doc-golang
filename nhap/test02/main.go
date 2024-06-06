package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
)

var IAM_URl = "https://staging.bizflycloud.vn/api/iam/v2"

var IAM_TOKEN = "7WNTmylN1EpWLgVwwI4JPi6GqGbZyQifOMoruXY9p7WkpmyuNlmzxlaYHMQv"

// dat.usth2000@gmail.com đã thực hiên GET https://hn-staging2.bizflycloud.vn/api/kubernetes-engine//_/e8d2vr30dfio7s94 tại 0ddd8351-bab2-4711-82c7-c53cf1e72b24 . Thời gian thực hiên Success action của user: 2024-03-18 17:03:48
type IAMGetOwnerInfoResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	ErrorCode int    `json:"error_code"`
	Data      struct {
		UUID  uuid.UUID `json:"uuid"`
		Email string    `json:"email"`
	} `json:"data"`
	ProjectDetail projectDetail `json:"project_detail"`
}
type projectDetail struct {
	AliasName string `json:"alias_name"`
}

func main() {

	url := fmt.Sprintf("%s/admin/project/owner_info?project_id=%s", IAM_URl, "0ddd8351-bab2-4711-82c7-c53cf1e72b24")
	//url := fmt.Sprintf("%s/projects/%s", IAM_URl, "048fcfc2-b338-4ca1-b1f9-2bc5234e3cae")
	res, err := CallAPI("GET", url, nil)
	if err != nil {
		fmt.Println("err ", err)
	}
	fmt.Println(string(res))
	a := IAMGetOwnerInfoResponse{}
	json.Unmarshal(res, &a)
	fmt.Println(a)
}
func CallAPI(method, url string, bodyJSON []byte) ([]byte, error) {
	// Tạo một client HTTP
	client := &http.Client{}

	// Tạo request đến API
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-IAM-TOKEN", IAM_TOKEN)

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
