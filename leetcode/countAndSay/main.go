package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(3322251))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := countAndSay(n - 1)
	ans := ""

	curr := s[0]
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == curr {
			count++
		} else {
			if count > 0 {
				ans += strconv.Itoa(count) + string(curr)
				count = 1
				curr = s[i]
			}
		}
	}
	if count > 0 {
		ans += strconv.Itoa(count) + string(curr)
		count = 0
		curr = byte(0)
	}
	return ans
}
