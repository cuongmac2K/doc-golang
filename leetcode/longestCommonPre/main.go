package main

import (
	"fmt"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("++", longestCommonPrefix(strs))
}
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	count := 1
	n := len(strs)
	char := []rune(strs[1])
	tem := string(char[0])
	arayR := make([]int, len(char))
	b := 0
	for i := 0; i < n-1; i++ {
		char1 := []rune(strs[i])
		char2 := []rune(strs[i+1])
		for j := 0; j < count; j++ {
			b++
			arayR[j] = 0
			if char1[j] == char2[j] && string(char1[j]) == tem {
				count++
				tem = string(char1[j])
				fmt.Println(tem)
				arayR[j]++
			}
			if b != count {
				continue
			}
		}

	}
	fmt.Println(arayR)
	return ""
}
