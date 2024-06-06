package main

import "fmt"

func main() {
	s := "cdf"
	d := "a"
	fmt.Println(mergeAlternately(s, d))
}

func mergeAlternately(word1 string, word2 string) string {
	n1 := len(word1)
	n2 := len(word2)
	s := ""
	if n1 == n2 {
		for i := 0; i < n2; i++ {
			s = s + string(word1[i]) + string(word2[i])
		}
	}
	if n1 > n2 {
		for i := 0; i < n2; i++ {
			s = s + string(word1[i]) + string(word2[i])
		}
		s = s + word1[n2:n1]
	}
	if n2 > n1 {
		for i := 0; i < n1; i++ {
			s = s + string(word1[i]) + string(word2[i])
		}
		s = s + word2[n1:n2]
	}
	return s
}
