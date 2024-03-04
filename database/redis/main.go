package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping(context.Background()).Result()
	fmt.Println(pong, err)

	//start1 := time.Now()
	//for i := 0; i < 10000; i++ {
	//	err = client.Set(context.Background(), "key"+strconv.Itoa(i), "xinchaongay moi  vui ve luon cuong nha", 6*time.Hour).Err()
	//	if err == nil {
	//		fmt.Println("ok")
	//	} else {
	//		fmt.Println("not ok: " + err.Error())
	//	}
	//	time.Sleep(time.Second / 100000)
	//}
	//fmt.Println(time.Since(start1))

	//del value
	start2 := time.Now()
	for i := 0; i < 10000; i++ {
		val, err := client.Del(context.Background(), "key"+strconv.Itoa(i)).Result()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("val", val)
		}
		time.Sleep(time.Second / 100000)
	}
	fmt.Println(time.Since(start2))

}
