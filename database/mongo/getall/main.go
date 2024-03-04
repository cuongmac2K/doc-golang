package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Record struct {
	ID   int    `bson:"_id"`
	Name string `bson:"name"`
}

func GetLastRecord(uri string) (*Record, error) {
	// Thiết lập kết nối MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Chọn cơ sở dữ liệu và bộ sưu tập
	db := client.Database("mydatabase")
	collection := db.Collection("mycollection")

	// Sắp xếp theo trường _id theo thứ tự giảm dần
	opts := options.FindOne().SetSort(bson.D{{"_id", -1}})

	// Tạo filter trống để lấy phần tử cuối cùng
	filter := bson.M{}

	// Tìm kiếm phần tử cuối cùng
	result := collection.FindOne(ctx, filter, opts)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var record Record
	err = result.Decode(&record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func main() {
	uri := "mongodb://root:123456@localhost:27017"
	record, err := GetLastRecord(uri)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Last record: %+v\n", record)
}
