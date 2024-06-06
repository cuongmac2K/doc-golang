//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/fluent/fluent-logger-golang/fluent"
//	"os"
//)
//
//var Flogger *fluent.Fluent
//
//func init() {
//	initFluent()
//
//}
//
//func initFluent() {
//	var err error
//	Flogger, err = fluent.New(fluent.Config{FluentPort: 24224, FluentHost: "127.0.0.1"})
//	if err != nil {
//		fmt.Printf("Could not connect to Fluent at %s Error : %v", os.Getenv("FLUENTHOST"), err)
//	}
//}
//
//// DebugLog (msg string)  Logs the string to Fluent server
//func DebugLog(msg string) {
//	//Flogger.Post("debug.access", map[string]string{"data": msg})
//	data := map[string]string{"data": msg}
//	jsonData, _ := json.Marshal(data)
//	err := Flogger.Post("debug.access", jsonData) // both methods give same result
//	fmt.Println("err: ", err)
//}
//
//func main() {
//	msg := `{"mykey":"myval"}`
//	fmt.Println(msg) //  prints {"mykey":"myval"}
//	DebugLog(msg)    //  prints {"data":"{\"mykey\":\"myval\"}"}   with extra \
//}

package main

import (
	"fmt"
	"github.com/fluent/fluent-logger-golang/fluent"
	"time"
)

func main() {
	logger, err := fluent.New(fluent.Config{FluentHost: "0.0.0.0", FluentPort: 8888})
	if err != nil {
		fmt.Println("Error creating Fluent logger:", err)
		return
	}
	defer logger.Close()

	tag := "fluentd-log-demo"
	data := map[string]interface{}{
		"Cuong": "John",
		"IdNo":  1234,
		"Ae":    25,
	}
	var a interface{}
	a = data
	// Gửi thông điệp với thời gian cụ thể
	err = logger.PostWithTime(tag, time.Now(), a)
	if err != nil {
		panic(err)
	}
}
