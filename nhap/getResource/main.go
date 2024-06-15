package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

func main() {

	//b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/iaas-cloud//servers/6034cd4b-73ca-4193-aff8-6b5722bd8a71",
	//	"GET", "98f5e866-ce84-41d8-b09c-d883a86e96c0")
	//b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/iaas-cloud/volume-types",
	//	"GET", "98f5e866-ce84-41d8-b09c-d883a86e96c0")

	//b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/kubernetes-engine/_/b8zdr8k5u4d8vbhx",
	//	"GET", "30f25e75-34c1-4ab8-80cc-0360dab126ba")

	//b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/iaas-cloud//servers/da311a20-cace-41d0-ab7d-8b1aace74c7e",
	//	"GET", "98f5e866-ce84-41d8-b09c-d883a86e96c0")

	//b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/auto-scaling/groups/591bbb04-9547-4412-afbb-eb8497067fc3",
	//	"DELETE", "1ee941ad-d7f5-4f00-8d51-42d29a6515dc")

	/*	b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/kubernetes-engine/_/m9mndvhbq3zacdvn",
		"DELETE", "30f25e75-34c1-4ab8-80cc-0360dab126ba")*/

	//b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/cloud-database/instances/289991bd-d5ba-458c-9e78-ad2975abc405",
	//	"DELETE", "a22f7b8c-037c-4ec9-86c8-f32f16546a80")

	//b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/iaas-cloud/vpc-networks/24781ad2-3c66-48da-8191-4c1739220971",
	//	"DELETE", "98f5e866-ce84-41d8-b09c-d883a86e96c0")

	//b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/iaas-cloud/keypairs/string_id_123ggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg",
	//	"GET", "98f5e866-ce84-41d8-b09c-d883a86e96c0")

	//b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/iaas-cloud/keypairs",
	//	"GET", "98f5e866-ce84-41d8-b09c-d883a86e96c0")

	//b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/domains/123e4567-e89b-42d3-a456-426614174000/namespaces/abjjc/get-capacity",
	//	"GET", "e733e886-0f8e-4929-82a7-a966d8189266")

	//b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/_/abcd1234efgh5678/0123456789ABCDEFabcdef012345/123e4567-e89b-42d3-a456-426614174000",
	//	"GET", "30f25e75-34c1-4ab8-80cc-0360dab126ba")
	//
	b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/kubernetes-engine/_/xvl083ux4ngewibb",
		"GET", "30f25e75-34c1-4ab8-80cc-0360dab126ba")

	//b := getValueResouce("https://hn-staging2.bizflycloud.vn/api/iaas-cloud/firewalls/2a0ce5b9-b05b-4741-a547-68eaa8d738dd",
	//	"GET", "98f5e866-ce84-41d8-b09c-d883a86e96c0")

	//b := getValueResouce(" https://hn-staging2.bizflycloud.vn/api/iaas-cloud/firewalls/2a0ce5b9-b05b-4741-a547-68eaa8d738dd",
	//	"Get", "98f5e866-ce84-41d8-b09c-d883a86e96c0")
	//b := getValueResouce("https://staging.bizflycloud.vn/api/al/visualize/explore/b0120bb8-ae50-4cbd-aab7-65835ec0fbd2?schemaVersion=1&panes=%20%20%20%20%20%20%7B%22mig%22:%7B%22datasource%22:%22f9301251-dd7e-4509-84b4-b1e2809e52b7%22,%22queries%22:%5B%7B%22refId%22:%22A%22,%22expr%22:%22%7B%20%20%20%20%20%20project_uuid%3D%5C%22b0120bb8-ae50-4cbd-aab7-65835ec0fbd2%5C%22%7D%20%7C%3D%20%60%60%22,%22queryType%22:%22range%22,%22datasource%22:%7B%22type%22:%22loki%22,%22uid%22:%22f9301251-dd7e-4509-84b4-b1e2809e52b7%22%7D,%22editorMode%22:%22builder%22%7D%5D,%22range%22:%7B%22from%22:%22now-1h%22,%22to%22:%22now%22%7D%7D%7D&orgId=1",
	//	"Get", "5338a9b6-7cd2-4b55-a2ff-aba26dea283e")
	fmt.Println("endpoint la: ", b)

}
func getValueResouce(resource string, method, service string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

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
	// sap xep mang dataResource theo url để check trường hợp url dài trước nếu ko thì các kết quả endpoint trả ra đều là kết quả của url ngắn
	sortDataResourcesByUrlLength(res1)

	if service == "98f5e866-ce84-41d8-b09c-d883a86e96c0" {
		for _, v := range res1 {
			if strings.ToUpper(v.Type) == strings.ToUpper(method) {
				if len(v.Patterns) == 0 {
					if strings.Contains(resource, v.Url) {
						//return v.Endpoint
						fmt.Println(" out url ", v.Url)
						fmt.Println("out record ", v)
						return v.Endpoint
					}
				} else {
					resouceCheck := resource
					var regexps []*regexp.Regexp
					for _, pattern := range v.Patterns {
						if pattern == "(?!\\/.)[A-Za-z0-9-_.@]{1,255}$" || pattern == "(?!\\/.)[A-Za-z0-9_.@]{1,255}$" {
							continue
						}

						regexps = append(regexps, regexp.MustCompile(pattern))
					}
					for _, re := range regexps {
						switch {
						case strings.Contains(re.String(), "(.*?)"):
							stringsRe := strings.Replace(re.String(), "(.*?)", "*", 1)
							resouceCheck = re.ReplaceAllString(resouceCheck, stringsRe)
						default:
							parts := strings.Split(resouceCheck, "/")
							for i, part := range parts {
								if re.MatchString(part) {
									parts[i] = "*"
									break
								}
							}
							resouceCheck = strings.Join(parts, "/")
						}
					}

					if strings.Contains(resouceCheck, v.Url) {
						//return v.Endpoint
						fmt.Println(" out url ", v.Url)
						fmt.Println("out record ", v)
						return v.Endpoint
					}
					// check case patterns is (?!\/.)[A-Za-z0-9_.@]{1,255}$
					resouceCheckKeyPairs, check := CheckPatternKeyPairs(resource)
					if check {
						if strings.Contains(resouceCheckKeyPairs, v.Url) {
							//return v.Endpoint
							fmt.Println(" out url ", v.Url)
							fmt.Println("out record ", v)
							return v.Endpoint
						}
					}
				}
			}
		}
	} else {
		for _, v := range res1 {
			if strings.ToUpper(v.Type) == strings.ToUpper(method) {
				if len(v.Patterns) == 0 {
					if strings.Contains(resource, v.Url) {
						//return v.Endpoint
						fmt.Println(" out url ", v.Url)
						fmt.Println("out record ", v)
					}
				} else {
					// Biên dịch các biểu thức chính quy
					var regexps []*regexp.Regexp
					for _, pattern := range v.Patterns {
						regexps = append(regexps, regexp.MustCompile(pattern))
					}
					resouceCheck := resource
					for _, re := range regexps {
						switch {
						case strings.Contains(re.String(), "(.*?)"):
							// Thay thế toàn bộ phần khớp nếu mẫu là namespaces/(.*?)/directories
							stringsRe := strings.Replace(re.String(), "(.*?)", "*", 1)
							resouceCheck = re.ReplaceAllString(resouceCheck, stringsRe)
						default:
							parts := strings.Split(resouceCheck, "/")
							for i, part := range parts {
								if re.MatchString(part) {
									parts[i] = "*"
									break
								}
							}
							resouceCheck = strings.Join(parts, "/")
						}
					}
					if strings.Contains(resouceCheck, v.Url) {
						//return v.Endpoint
						fmt.Println(" out url ", v.Url)
						fmt.Println("out record ", v)
						return v.Endpoint
					}
				}

			}
		}
	}
	return resource
}

type dataResource struct {
	Patterns []string `json:"patterns"`
	Type     string   `json:"type"`
	Endpoint string   `json:"endpoint"`
	Url      string   `json:"url"`
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
func CheckPatternKeyPairs(url string) (string, bool) {
	if !strings.Contains(url, "/keypairs/") {
		return "", false
	}
	// vì trong bảng resource_mapping có 2 record đều có url keypairs nên phải check 2 patterns

	// Định nghĩa regex pattern thứ nhất
	pattern1 := `^[A-Za-z0-9_.@]{1,255}$`
	re1 := regexp.MustCompile(pattern1)

	// Định nghĩa regex pattern thứ hai
	pattern2 := `^[A-Za-z0-9-_.@]{1,255}$`
	re2 := regexp.MustCompile(pattern2)

	// Tìm vị trí cuối cùng của dấu "/"
	lastSlashIndex := strings.LastIndex(url, "/")

	if lastSlashIndex != -1 {
		// Tách phần cuối cùng của URL
		lastPart := url[lastSlashIndex+1:]

		// Kiểm tra nếu phần cuối không bắt đầu với "/"
		if !strings.HasPrefix(lastPart, "/") && (re1.MatchString(lastPart) || re2.MatchString(lastPart)) {
			// Thay thế phần cuối bằng "*"
			url = url[:lastSlashIndex+1] + "*"
		}
	}

	return url, true
}

// Hàm để sắp xếp mảng dataResource dựa trên độ dài của Url
func sortDataResourcesByUrlLength(resources []dataResource) {
	sort.Slice(resources, func(i, j int) bool {
		return len(resources[i].Url) > len(resources[j].Url)
	})
}
