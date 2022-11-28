package main

import "fmt"

func main() {
	//s := [][]string{{"lam sao"}, {"gion ko"}}
	//s[0][0] = "cuong"
	//fmt.Println(s)
	s := "babdff"
	checkDoiXung(s)
	fmt.Println("len ", len(s))
}

func checkDoiXung(s string) bool {
	chars1 := []rune(s)
	len1 := len(chars1)
	for i := 0; i < len1; i++ {
		if chars1[i] != chars1[len1-i-1] {
			return false
		}
	}
	return true
}
