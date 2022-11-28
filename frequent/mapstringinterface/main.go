package main

import "fmt"

func main() {
	//secretString := "{\"type\":\"string\",\"value\":\"HN\"}"
	//
	//sec := map[string]interface{}{}
	//if err := json.Unmarshal([]byte(secretString), &sec); err != nil {
	//	panic(err)
	//}
	//s, _ := sec["value"].(string)
	//fmt.Println(s)
	count := map[int]int{}
	i, ok := count[0]
	fmt.Println(i, ok)
	count[0] = 4

	i1, ok1 := count[0]
	fmt.Println(i1, ok1)
	fmt.Println(count)
}
