package main

// Import thêm gói ioutil
import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/writeLog1", writeLogHandler1).Methods("POST")
	r.HandleFunc("/readLog1", readLogHandler1).Methods("GET")

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3000", r))
}

func writeLogHandler1(w http.ResponseWriter, r *http.Request) {
	// Đọc dữ liệu từ request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Ghi dữ liệu log từ request body vào file log
	logMessage := string(body)
	log.Println(logMessage)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Log has been written to the file"))
}

func readLogHandler1(w http.ResponseWriter, r *http.Request) {
	// Đọc nội dung log từ tập tin và trả về cho client
	logFile, err := os.Open("app.log")
	if err != nil {
		http.Error(w, "Failed to read log file", http.StatusInternalServerError)
		return
	}
	defer logFile.Close()

	info, err := logFile.Stat()
	if err != nil {
		http.Error(w, "Failed to read log file", http.StatusInternalServerError)
		return
	}

	data := make([]byte, info.Size())
	_, err = logFile.Read(data)
	if err != nil {
		http.Error(w, "Failed to read log file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
