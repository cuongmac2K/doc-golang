package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Record struct {
	ID   int    `bson:"_id"`
	Name string `bson:"name"`
}

func main() {
	// Thiết lập kết nối MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
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

	// Tạo một slice chứa tỷ lệ record
	records := make([]interface{}, 0, 1000000000)

	// Tạo tỷ lệ record
	for i := 1; i <= 1000000000; i++ {
		record := Record{
			ID:   i,
			Name: fmt.Sprintf("Record %d", i),
		}

		records = append(records, record)
	}

	// Chia nhỏ tỷ lệ record thành các batch nhỏ để chèn vào MongoDB
	batchSize := 1000
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

		fmt.Printf("Inserted batch %d/%d\n", i+1, totalBatches)
	}

	fmt.Println("Insertion completed")
}
