package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getFormattedDate() string {
	currentTime := time.Now()
	dateStr := currentTime.Format("02/01")
	return dateStr
}
func main() {
	fmt.Println(getFormattedDate())
	http.HandleFunc("/api/test-post", getUser)
	http.ListenAndServe(":4000", nil)
}

func getUser(writer http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)
	data := map[string]interface{}{}
	json.Unmarshal(body, &data)
	name := data["name"]

	aa := map[string]interface{}{
		"message": "success",
		"name":    name,
	}
	byte_res, _ := json.Marshal(aa)
	writer.Write(byte_res)
	writer.WriteHeader(200)
}
