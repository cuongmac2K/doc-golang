package main

import "fmt"

func main() {
	haystack := "avfvfbgbhnhnhnbc"
	needle := "c"
	fmt.Println(strStr(haystack, needle))
}
func strStr(haystack string, needle string) int {
	n := len(haystack)
	if n == 1 {
		return 0
	}
	for i := 0; i <= n; i++ {
		for j := i; j <= n; j++ {
			s := haystack[i:j]
			fmt.Println(s)
			if s == needle {
				return i
			}
		}
	}
	return -1
}
