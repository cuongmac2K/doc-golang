package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type StreamData struct {
	Project     string `json:"project"`
	RequestId   string `json:"request_id"`
	ProjectUuid string `json:"project_uuid"`
	Service     string `json:"service"`
	User        string `json:"user"`
	Category    string `json:"category"`
}
type Stream struct {
	Stream StreamData      `json:"stream"`
	Values [][]interface{} `json:"values"`
}
type Payload struct {
	Streams []Stream `json:"streams"`
}

type AgentData struct {
	Project     string `json:"project"`
	RequestId   string `json:"request_id"`
	ProjectUuid string `json:"project_uuid"`
	Service     string `json:"service"`
	User        string `json:"user"`
	Category    string `json:"category"`
}

func main() {

	for i := 0; i < 10; i++ {
		for i := 0; i < 10; i++ {
			agentData := AgentData{
				Project:     "cuongmac2k@gmail.com",
				RequestId:   "req-23fb53b7-3de8-4cb2-ad0a-11be58daf82b",
				ProjectUuid: "b0120bb8-ae50-4cbd-aab7-65835ec0fbd2",
				Service:     "IAM",
				User:        "cuongmac2k@gmail.com",
				Category:    "AUDITLOG",
			}
			res := "daylatktest2@bizflycloud.vn đã thực hiên GET https://hn-staging2.bizflycloud.vn/api/iaas-cloud/flavors tại daylatktest2@bizflycloud.vn . Thời gian thực hiên Success action của user: 2024-03-19 10:23:02"

			PushDataToLoKi(agentData, res)
		}
	}
}
func PushDataToLoKi(agentData AgentData, logs string) {
	//url := fmt.Sprintf("%s/loki/api/v1/push", "http://10.3.93.217:30266")
	//- CLICKHOUSE_SERVER=10.3.53.206
	//- CLICKHOUSE_AUTH=default:admin
	//- CLICKHOUSE_DB=qryn2
	url := fmt.Sprintf("%s/loki/api/v1/push", "http://0.0.0.0:3200")
	payload := Payload{
		Streams: []Stream{
			{
				Stream: StreamData{
					Project:     agentData.Project,
					RequestId:   agentData.RequestId,
					ProjectUuid: agentData.ProjectUuid,
					Service:     agentData.Service,
					User:        agentData.User,
					Category:    agentData.Category,
				},
				Values: [][]interface{}{
					{
						time.Now().UnixNano(),
						logs,
					},
				},
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling payload:", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Set("Content-Type", "application/json")
	username := "admin"
	password := "ADMINn0taAEmin@13579"
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}
	fmt.Println(string(body))
}
