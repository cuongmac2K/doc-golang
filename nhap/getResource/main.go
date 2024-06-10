package main

import (
	"bytes"
	"doc-golang/nhap/logger"
	"encoding/json"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func main() {
	a := time.Now()
	b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/iaas-cloud/firewalls/2a0ce5b9-b05b-4741-a547-68eaa8d738dd",
		"PATCH", "98f5e866-ce84-41d8-b09c-d883a86e96c0")
	//b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/kubernetes-engine/_/b8sm4tbo4lg2frb2", "Get")
	//b := getValueResouce("https://staging.bizflycloud.vn/api/al/visualize/explore/b0120bb8-ae50-4cbd-aab7-65835ec0fbd2?schemaVersion=1&panes=%20%20%20%20%20%20%7B%22mig%22:%7B%22datasource%22:%22f9301251-dd7e-4509-84b4-b1e2809e52b7%22,%22queries%22:%5B%7B%22refId%22:%22A%22,%22expr%22:%22%7B%20%20%20%20%20%20project_uuid%3D%5C%22b0120bb8-ae50-4cbd-aab7-65835ec0fbd2%5C%22%7D%20%7C%3D%20%60%60%22,%22queryType%22:%22range%22,%22datasource%22:%7B%22type%22:%22loki%22,%22uid%22:%22f9301251-dd7e-4509-84b4-b1e2809e52b7%22%7D,%22editorMode%22:%22builder%22%7D%5D,%22range%22:%7B%22from%22:%22now-1h%22,%22to%22:%22now%22%7D%7D%7D&orgId=1", "Get")
	fmt.Println(b)
	fmt.Println(time.Since(a))
}
func getValueResouce(resource string, method, service string) string {
	url := fmt.Sprintf("%s/admin/resources?%s%s", "https://staging.bizflycloud.vn/api/iam/v2", "service_uuid=", service)
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
	for _, v := range res1 {
		if strings.ToUpper(v.Type) == strings.ToUpper(method) {
			if len(v.Patterns) > 0 {
				s := resource
				for _, i := range v.Patterns {
					re := regexp.MustCompile(i)

					// Thay thế tất cả các chuỗi khớp với regex bằng dấu *
					result := re.ReplaceAllString(s, "*")
					//fmt.Println("reslut sau khi thay *: ", result)
					//
					if strings.Contains(result, v.Url) {
						//return v.Endpoint
						fmt.Println("url ", v.Url)
						fmt.Println("enpoint ", v.Endpoint)
						//fmt.Println("record ", v)
						break
					}
				}

			}
			if strings.Contains(resource, v.Url) {
				//return v.Endpoint
				fmt.Println(" out url ", v.Url)
				fmt.Println("out record ", v)
			}

		}
	}
	return ""
}

type dataResource struct {
	Patterns    []string `json:"patterns"`
	Type        string   `json:"type"`
	Endpoint    string   `json:"endpoint"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
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
func getValueProject(project_uuid string) string {
	url := fmt.Sprintf("%s/admin/project/owner_info?project_id=%s", "https://staging.bizflycloud.vn/api/iam/v2", project_uuid)
	body, err := CallAPI("GET", url, nil)
	if err != nil {
		logger.Debugf("err call get nameProject: %v", err)
		sentry.CaptureException(err)
		return ""
	}
	res := IAMGetOwnerInfoResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		logger.Debugf("err unmarshal value get nameProject: %v", err)
		sentry.CaptureException(err)
		return ""
	}
	return res.Data.ProjectDetail.AliasName
}

type IAMGetOwnerInfoResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	ErrorCode int    `json:"error_code"`
	Data      struct {
		UUID          uuid.UUID     `json:"uuid"`
		Email         string        `json:"email"`
		ProjectDetail projectDetail `json:"project_detail"`
	} `json:"data"`
}
type projectDetail struct {
	AliasName string `json:"alias_name"`
}
