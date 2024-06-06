package main

import (
	"encoding/json"
	"log"
	"net"
	"time"
)

func main() {
	message := map[string]interface{}{
		"title":        "EXTRACT DATA FOR AUDIT LOGS",
		"request_id":   "req-fd8fea44-88ad-9260-b6a3-e150cbeb2396",
		"ipv4":         "123.31.45.180",
		"user_uuid":    "7ebdbc6e-98db-4681-aaef-5e18b974fda1",
		"owner_uuid":   "7ebdbc6e-98db-4681-aaef-5e18b974fda1",
		"email":        "missyou97s2@gmail.com",
		"action":       "GET",
		"resource":     "https://hn-staging2.bizflycloud.vn/api/auto-scaling/task/d226bcb5-da2d-4ec2-8569-8a90e1f8eb5f/status",
		"service":      "1ee941ad-d7f5-4f00-8d51-42d29a6515dc",
		"status_code":  200,
		"error":        "",
		"project_uuid": "2c50ce8d-e598-44ff-8249-932ef2becbbc",
		"request_time": time.Now().Format("2006-01-02 15:04:05"),
		"header":       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb29raWUiOiJ7XCI6YXV0aG9yaXR5XCI6W1wiMTAuNS43MC4yMDE6MzAzMzlcIl0sXCI6bWV0aG9kXCI6W1wiR0VUXCJdLFwiOnBhdGhcIjpbXCIvdGFzay9kMjI2YmNiNS1kYTJkLTRlYzItODU2OS04YTkwZTFmOGViNWYvc3RhdHVzXCJdLFwiOnNjaGVtZVwiOltcImh0dHBzXCJdLFwiQWNjZXB0XCI6W1wiKi8qXCJdLFwiQWNjZXB0LUVuY29kaW5nXCI6W1wiZ3ppcCwgZGVmbGF0ZVwiXSxcIkNvbnRlbnQtTGVuZ3RoXCI6W1wiMlwiXSxcIkNvbnRlbnQtVHlwZVwiOltcImFwcGxpY2F0aW9uL2pzb25cIl0sXCJVc2VyLUFnZW50XCI6W1wicHl0aG9uLXJlcXVlc3RzLzIuMjcuMVwiXSxcIlgtQXV0aC1Ub2tlblwiOltcImdBQUFBQUJtUXgxVElXWjNNWUpDQXpXNzZwUlV3Q3VocktfQTlXOE5ZMDNPYzU2djN2LXFNQ3oteWlUazBfajVIQ3Q2SXFjeHBwQzlfalBNb0RhWFNNNV9sa3JONFV0alZMTWlNWURHUkRlSWNta3lJX1FNVGEwTjVMV2pZX3k5UHROdmFZZmtidlpESHhRSUJyN19aendTQmZmZUdxSlRSd0dmQkEwek9vcWZoZnJ4Vlh4R1R",
	}
	sendLog(message)
}
func sendLog(message map[string]interface{}) {
	conn, err := net.Dial("tcp", "0.0.0.0:24224")
	if err != nil {
		log.Fatalf("Failed to connect to Fluent Bit: %v", err)
	}
	defer conn.Close()
	logData := map[string]interface{}{
		"message":              message,
		"decrypt_data_service": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjYXRlZ29yeSI6IkFVRElUTE9HIiwiZW1haWwiOiJjdW9uZ21hYzJrQG1haWwuY29tIiwibmFtZSI6ImFnZW50IDU1NTU1NTU1NTU1NTU1NTU1IiwidG9rZW5faW5fd29ya2VyIjoiOWE2NDMyZTAtOGRkZi00NmFkLTlmY2MtYjg4ZWM4ODhkM2RiIn0.FBH9JDxdCRBBFsi8prZJSDj5L1wCH4Ca7K1ZrQ_M2bA",
		"secret":               "abc",
	}
	jsonData, err := json.Marshal(logData)
	if err != nil {
		log.Fatalf("Failed to marshal log data to JSON: %v", err)
	}

	_, err = conn.Write(jsonData)
	if err != nil {
		log.Fatalf("Failed to send log: %v", err)
	}

	log.Println("Log sent successfully.")
}
