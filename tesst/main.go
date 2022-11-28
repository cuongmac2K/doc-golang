package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	query_update := map[string]interface{}{
		"script": map[string]interface{}{
			"source": "source",
			"lang":   "painless",
		},
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"user_sequence_id": "dffffg",
			},
		},
	}
	convert(query_update)

	//queryString := fmt.Sprintf("%v", query_update)
	//fmt.Println(queryString)
	//d := []byte{}
	//d, _ = json.Marshal(query_update)
	//fmt.Println(string(d))
	//fmt.Println("cuong bede")
	//
	fmt.Println(4|5, "=========\n")

	filter := bson.M{
		"id_crm":  456,
		"project": 666,
		"oaId":    545,
	}
	//d := []byte{}
	//d, _ = json.Marshal(filter)
	fmt.Println(filter)
}
func convert(record interface{}) {
	fmt.Println(record)
}
