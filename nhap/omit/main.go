package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
}

type Person struct {
	Name    string  `json:"name,omitempty"`
	Age     int     `json:"age,omitempty"`
	Email   string  `json:"email,omitempty"`
	Address Address `json:"address,omitempty"`
}

func main() {
	p1 := Person{Name: "Alice", Age: 30, Email: "alice@example.com", Address: Address{City: "New York", Country: "USA"}}
	p2 := Person{Name: "Bob", Age: 0, Email: "", Address: Address{}}

	// Convert structs to JSON
	jsonStr1, _ := json.Marshal(p1)
	jsonStr2, _ := json.Marshal(p2)

	fmt.Println(string(jsonStr1)) // {"name":"Alice","age":30,"email":"alice@example.com","address":{"city":"New York","country":"USA"}}
	fmt.Println(string(jsonStr2)) // {"name":"Bob"}
	pp := Person{}
	_ = json.Unmarshal(jsonStr1, &pp)
	ppp := Person{}
	_ = json.Unmarshal(jsonStr2, &ppp)
	fmt.Println(pp)
	fmt.Println(ppp)
}
