package main

import (
	"fmt"
)

func main() {
	s := "pwwkew"
	fmt.Println("out put ", lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	var start, end, output int
	var visited = make(map[rune]int)
	for i, c := range s {
		end = i
		if start1, found := visited[c]; found {
			if start1+1 > start {
				start = start1 + 1
			}
		}
		if output < end-start+1 {
			output++
		}
		visited[c] = end
	}
	return output
}
