package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.OpenFile("/var/log/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	logger := log.New(file, "", log.LstdFlags)

	for i := 0; i < 20; i++ {
		if err != nil {
			log.Fatal(err)
		}

		logger.Println("log to agent 1 { " + strconv.Itoa(i))
		//time.Sleep(time.Second / 2)
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
