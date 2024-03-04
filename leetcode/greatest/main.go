package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "TAUXXTAUXXTAUXXTAUXXTAUXX"
	s2 := "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"
	s3 := s1 + s2
	fmt.Println(gcdOfStrings(s3, s2))
}
func gcdOfStrings(str1 string, str2 string) string {
	n := len(str2)
	for i := 0; i < n; i++ {
		for j := n; j > i+1; j-- {
			a := strings.Replace(str1, str2[i:j], "", -1)
			if a == "" {
				return str2[i:j]
			}
		}

	}
	return ""
}
