package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func main() {
	getUser()
}

func getUser() {
	s := `{"date":1711006697.228178,"stream":"stdout","time":"2024-03-21T07:38:17","source":"edge-proxy","severity":"INFO","message":{"title":"EXTRACT DATA FOR AUDIT LOGS","request_id":"req-36c3e28f-b427-976e-bbb2-2b29b8f176bf","ipv4":"123.31.45.180","user_uuid":"19a55f97-394a-4055-b83f-da55892e17d8","owner_uuid":"19a55f97-394a-4055-b83f-da55892e17d8","email":"nguyenphuongthao04@vccorp.vn","action":"GET","resource":"https://hn-staging2.bizflycloud.vn/api/kubernetes-engine/11/_/node_everywhere/d2ba672e-4fd8-4b17-9777-a4d8aafae6f4","service":"30f25e75-34c1-4ab8-80cc-0360dab126ba","status_code":200,"error":"","project_uuid":"9d8ccb77-cc21-4b06-a063-ed1d2ef2bf70","request_time":"2024-03-21 14:38:17","header":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb29raWUiOiJ7XCI6YXV0aG9yaXR5XCI6W1wiZW5naW5lLmFwaS5rOHNhYXMtc3RhZ2luZy5iaXpmbHljbG91ZC52blwiXSxcIjptZXRob2RcIjpbXCJHRVRcIl0sXCI6cGF0aFwiOltcIi9fL25vZGVfZXZlcnl3aGVyZS9kMmJhNjcyZS00ZmQ4LTRiMTctOTc3Ny1hNGQ4YWFmYWU2ZjRcIl0sXCI6c2NoZW1lXCI6W1wiaHR0cFwiXSxcIkFjY2VwdFwiOltcImFwcGxpY2F0aW9uL2pzb247IGNoYXJzZXQ9dXRmLThcIl0sXCJBY2NlcHQtRW5jb2RpbmdcIjpbXCJnemlwXCJdLFwiQXV0aG9yaXphdGlvblwiOltcIkJhc2ljXCJdLFwiQ29udGVudC1UeXBlXCI6W1wiYXBwbGljYXRpb24vanNvbjsgY2hhcnNldD11dGYtOFwiXSxcIlVzZXItQWdlbnRcIjpbXCJiaXpmbHktY2xpZW50LWdvLzAuMC4xXCJdLFwiWC1BdXRoLVRva2VuXCI6W1wiZ0FBQUFBQmwtOC1vVjN0Y3VaWkd4UUFQNzRIMTFCR0FBc0ZpWDlIUGVHcElkTzlFbjFWY1ZTVF96WUxaV21WVTdqcE5QMjgtRHNoZ2NDSFl6c0R4RUxTNVg4U0V1N3F1V0ZwQVRZVDVoUmJlY1lPVFZtOXJ6UjNITThIazZVRkhmak9hcV93bkhVYjRqSHZSZi1xNnhlUkI0RWNrb1hnQ1V0bm9YVEpOY1UxWHJ1YUtUSUxXU0x3TGNMQjJDM21RRGZxc0VNOVZaeWozXCJdLFwiWC1BdXRoLVR5cGVcIjpbXCJ0b2tlblwiXSxcIlgtRW52b3ktRXh0ZXJuYWwtQWRkcmVzc1wiOltcIjEwLjI0NC4xLjJcIl0sXCJYLUZvcndhcmRlZC1Gb3JcIjpbXCIxMjMuMzEuMTEuMTM1LDEwLjI0NC4xLjJcIl0sXCJYLUZvcndhcmRlZC1Ib3N0XCI6W1wiaG4tc3RhZ2luZzIuYml6Zmx5Y2xvdWQudm5cIl0sXCJYLUZvcndhcmRlZC1Qcm90b1wiOltcImh0dHBcIl0sXCJYLU9yaWdpbi1Ib3N0XCI6W1wiaG4tc3RhZ2luZzIuYml6Zmx5Y2xvdWQudm5cIl0sXCJYLU9yaWdpbi1QYXRoXCI6W1wiL2FwaS9rdWJlcm5ldGVzLWVuZ2luZS8vXy9ub2RlX2V2ZXJ5d2hlcmUvZDJiYTY3MmUtNGZkOC00YjE3LTk3NzctYTRkOGFhZmFlNmY0XCJdLFwiWC1PcmlnaW5hbC1Gb3J3YXJkZWQtRm9yXCI6W1wiMTIzLjMxLjQ1LjE4MFwiXSxcIlgtUHJvamVjdC1JZFwiOltcIjlkOGNjYjc3Y2MyMTRiMDZhMDYzZWQxZDJlZjJiZjcwXCJdLFwiWC1SZWFsLUlwXCI6W1wiMTIzLjMxLjQ1LjE4MFwiXSxcIlgtUmVxdWVzdC1JZFwiOltcIjM2YzNlMjhmLWI0MjctOTc2ZS1iYmIyLTJiMjliOGYxNzZiZlwiXX0ifQ._7q1-HXKFP7xLAiXbu1r0TviUsbfDhH_f92UJ5HKUkM","cookie":"","payload":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb29raWUiOiIifQ.4qdnbA8n0zNRj35fn4lNgVRvmdUjlLNU8JYCEU2FhTw"},"kubernetes":{"pod_name":"edge-proxy-7d6f68866f-2vn89","namespace_name":"staging","pod_id":"69c2f264-4a0f-4b02-a97f-c144d977f591","labels":{"app.kubernetes.io/instance":"edge-proxy","app.kubernetes.io/name":"edge-proxy","pod-template-hash":"7d6f68866f"},"annotations":{"kubectl.kubernetes.io/restartedAt":"2023-08-29T02:12:34Z"},"host":"kubernetes-hn-staging-worker-node-10","container_name":"envoy-bussines-ext","docker_id":"8489acf89c41d9694c0ea19c0baabef071dbf8cb027283fe72ed55ff4795034f","container_hash":"hub.paas.vn/openstack/edge-proxy@sha256:70c09f3a828903484127b634d6786eef1403d0ef19c0f4d35267fdfd83f90469","container_image":"hub.paas.vn/openstack/edge-proxy:s160069_83c981a8"}}`
	for i := 0; i < 10; i++ {
		s += s
	}
	byte, _ := json.Marshal(s)
	CallAPILog("POST", "http://0.0.0.0:6060/logs/create", byte, "abc", "abc", "abc")
}
func CallAPILog(method, url string, bodyJSON []byte, xToken, project, uuid string) ([]byte, error) {
	// Tạo một client HTTP
	client := &http.Client{}

	// Tạo request đến API
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X_Token", xToken)
	req.Header.Set("Project_Id", project)
	req.Header.Set("Uuid", uuid)
	req.Header.Set("Content-Type", "application/json")
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
