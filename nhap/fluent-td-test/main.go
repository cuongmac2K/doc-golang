package main

import (
	"encoding/json"
	"log"
	"net"
)

func main() {
	message := map[string]interface{}{
		"name_log": "test log",
		"address":  "Ha Noi",
	}
	sendLog(message)
}

func sendLog(message map[string]interface{}) {
	conn, err := net.Dial("tcp", "0.0.0.0:24225")
	if err != nil {
		log.Fatalf("Failed to connect to Fluent Bit: %v", err)
	}
	defer conn.Close()

	logData := map[string]interface{}{
		"log":                  message,
		"decrypt_data_service": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjYXRlZ29yeSI6Ik5ldyIsIm5hbWUiOiJzZGsgNTY2NjY2IiwicHJvamVjdF9pZCI6ImIwMTIwYmI4LWFlNTAtNGNiZC1hYWI3LTY1ODM1ZWMwZmJkMiIsInRva2VuX2luX3dvcmtlciI6IiIsInR5cGUiOiJzZGsiLCJ1c2VyIjoiY3VvbmdtYWMya0BtYWlsLmNvbSJ9.LRfRILawil2sSJwiQ7ELBKNLXsrGFokyLM8kTMiinsM",
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
