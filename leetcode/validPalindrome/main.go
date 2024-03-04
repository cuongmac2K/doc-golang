package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := "A man, a plan, a canal: Panama"
	s2 := "aba"
	fmt.Println(isPalindrome(s1))
	fmt.Println(isPalindrome(s2))
}
func isPalindrome(s string) bool {
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}

		return unicode.ToLower(r)
	}
	s = strings.Map(f, s)

	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}

		i = i + 1
		j = j - 1
	}

	return true
}
