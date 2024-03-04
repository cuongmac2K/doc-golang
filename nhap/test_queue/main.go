package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/hello1", hello1)
	http.HandleFunc("/hello3", hello3)
	http.HandleFunc("/hello2", hello2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func hello1(w http.ResponseWriter, r *http.Request) {
	//newNameQueue := "name-queue-" + strconv.Itoa(1)
	//newQueue := []map[string]interface{}{
	//	{
	//		"queueName":   newNameQueue,
	//		"rateLimit":   2,
	//		"stopQueue":   false,
	//		"shutdown":    false,
	//		"waitingTask": true},
	//}
	//jsonData, _ := json.Marshal(newQueue)
	//_, _ = SetupNewQueue(
	//	"",
	//	r,
	//	jsonData,
	//)

	_, _ = DeleteQueue("", r, "name-queue-"+strconv.Itoa(80))

	//for i := 0; i < 10; i++ {
	//	_, _ = AddNewTaskQueue(
	//		"",
	//		r,
	//		"/hello2",
	//		url.Values{"name": {strconv.Itoa(i)}},
	//		"3",
	//	)
	//}

	_, _ = w.Write([]byte("day la hello 1"))
}

func hello3(w http.ResponseWriter, r *http.Request) {
	for j := 1; j < 81; j++ {
		newNameQueue := "name-queue-" + strconv.Itoa(j)
		newQueue := []map[string]interface{}{
			{
				"queueName":   newNameQueue,
				"rateLimit":   1,
				"stopQueue":   false,
				"shutdown":    false,
				"waitingTask": true},
		}
		jsonData, _ := json.Marshal(newQueue)
		_, _ = SetupNewQueue(
			"",
			r,
			jsonData,
		)
		time.Sleep(1 / 50)
		for i := 0; i < 70; i++ {
			_, _ = AddNewTaskQueue(
				"",
				r,
				"/hello2",
				url.Values{"name": {strconv.Itoa(j)}},
				newNameQueue,
			)
		}

	}
	_, _ = w.Write([]byte("day la hello 1"))
}

func hello2(w http.ResponseWriter, r *http.Request) {

	fmt.Println("da goi vao hello 2 and j = ", r.FormValue("name"))
	//_, _ = w.Write([]byte("day la hello 2"))
}
