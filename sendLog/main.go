package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		file, err := os.OpenFile("/var/log/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		logger := log.New(file, "", log.LstdFlags)
		logger.Println("log to agent 1 " + strconv.Itoa(i))
		time.Sleep(time.Second / 3)
		//file1, err := os.OpenFile("/var/log/test1.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//defer file1.Close()
		//
		//logger1 := log.New(file1, "", log.LstdFlags)
		//logger1.Println("log to agent 2", strconv.Itoa(i))
	}
	fmt.Println("finish")
}
