package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

type Record struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func main() {
	//for i := 0; i < 100; i++ {
	// Thiết lập kết nối MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:123456@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Chọn cơ sở dữ liệu và bộ sưu tập
	db := client.Database("mydatabase")
	collection := db.Collection("mycollection")

	//for i := 0; i < 10; i++ {
	// Tạo một slice chứa tỷ lệ record
	records := make([]interface{}, 0, 1000000)

	// Tạo tỷ lệ record
	for i := 1; i <= 1000000; i++ {
		record := Record{
			ID:   strconv.Itoa(i*10000000000+time.Now().UTC().Nanosecond()) + "h",
			Name: fmt.Sprintf("Record %d", i),
		}
		//fmt.Println(record)
		records = append(records, record)
	}

	// Chia nhỏ tỷ lệ record thành các batch nhỏ để chèn vào MongoDB
	batchSize := 10000
	totalBatches := len(records) / batchSize

	for i := 0; i < totalBatches; i++ {
		startIndex := i * batchSize
		endIndex := (i + 1) * batchSize

		batch := records[startIndex:endIndex]

		// Chèn batch vào MongoDB
		_, err := collection.InsertMany(ctx, batch)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Millisecond * 10)
		//	fmt.Printf("Inserted batch %d/%d\n", i+1, totalBatches)
	}
	//}
	//	time.Sleep(time.Second / 5)
	//}
	fmt.Println("Insertion completed")
}
