package main

import (
	"encoding/json"
	"log"
	"net"
)

// Hàm main để bắt đầu chương trình
func main() {
	message := map[string]interface{}{
		"name_log": "xin chao xin chao ",
		"address":  "Ha Noi",
	}
	sendLog(message)
}

// Hàm sendLog để gửi log tới Fluent Bit
func sendLog(message map[string]interface{}) {
	conn, err := net.Dial("tcp", "10.3.53.55:24224")
	if err != nil {
		log.Fatalf("Failed to connect to Fluent Bit: %v", err)
	}
	defer conn.Close()

	logData := map[string]interface{}{
		"log":                  message,
		"decrypt_data_service": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjYXRlZ29yeSI6Ik5ldyIsIm5hbWUiOiJzZGsgNTY2NjY2IiwicHJvamVjdF9pZCI6ImIwMTIwYmI4LWFlNTAtNGNiZC1hYWI3LTY1ODM1ZWMwZmJkMiIsInRva2VuX2luX3dvcmtlciI6IjlhNjQzMmUwLThkZGYtNDZhZC05ZmNjLWI4OGVjODg4ZDNkYiIsInR5cGUiOiJzZGsiLCJ1c2VyIjoiY3VvbmdtYWMya0BtYWlsLmNvbSJ9.2cObFwn_i3W-dC1yAoHCkxZNboJPexh47YeVtVrbDsQ",
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
