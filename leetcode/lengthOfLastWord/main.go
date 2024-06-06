package main

import (
	"fmt"
)

func main() {
	//s := "   fly me   to   the moon  "
	s := "Hello World"
	//s := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))
}
func lengthOfLastWord(s string) int {
	n := len(s)
	k := 0
	for i := n - 1; i > 1; i-- {
		if s[i] != 32 && s[i-1] != 32 {
			k = k + 1
		}
		if s[i] != 32 && s[i-1] == 32 {
			break
		}
	}
	return k + 1
}
