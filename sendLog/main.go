package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("/var/log/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	logger := log.New(file, "", log.LstdFlags)

	for i := 0; i < 2000; i++ {
		if err != nil {
			log.Fatal(err)
		}
		logger.Println("[abc@gmail.com]-[tcqshinns-MacBook-Pro]-[syslogd[115]: ASL Sender Statistics]-[10.3.251.144]-[Ha-Noi]")

		//time.Sleep(time.Millisecond / 2)
		//file1, err := os.OpenFile("/var/log/test1.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		//if err != nil {
		//	log.Fatal(err
		//}
		//defer file1.Close()
		//
		//logger1 := log.New(file1, "", log.LstdFlags)
		//logger1.Println("log to agent 2", strconv.Itoa(i))
	}
	fmt.Println("finish")
}
