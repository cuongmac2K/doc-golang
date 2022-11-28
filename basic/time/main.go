package main

import (
	"fmt"
	"strings"
)

func main() {
	//now := time.Now()
	//fmt.Println(now.Format(" 15:04 2006-01-02"))
	//
	//cur := "2020-03-10T08:21:29.435Z"
	//fmt.Println(TimeTrue(cur))
	a := fmt.Sprintf("abc la %s", "dfff")
	fmt.Println("aa ", a)
}
func TimeTrue(s string) string {
	var result string
	a := strings.Split(s, "T")
	b := a[1][:5]
	c := strings.Split(a[0], "-")
	result = b + " " + c[2] + "-" + c[1] + "-" + c[0]
	return result
}
