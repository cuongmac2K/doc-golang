package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "   the    sky is blue"
	fmt.Println(reverseWords(s))
}
func reverseWords(s string) string {
	words := strings.Fields(s)
	ar := []string{}
	n := len(words)
	for i := n - 1; i >= 0; i-- {
		a := strings.TrimSpace(words[i])
		ar = append(ar, a)
	}
	return strings.Join(ar, " ")
}
