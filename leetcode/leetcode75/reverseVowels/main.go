package main

import "fmt"

func main() {
	s := "leetcode"
	fmt.Println(reverseVowels(s))
}
func reverseVowels(s string) string {
	right := len(s) - 1
	left := 0
	a := []byte(s)

	for right > left {
		for !isVowel(a[left]) && left < len(s)-1 {
			left++
		}

		for !isVowel(a[right]) && right > 0 {
			right--
		}
		if right > left {
			a[left], a[right] = a[right], a[left]
			right--
			left++
		}
	}
	return string(a)

}

func isVowel(b byte) bool {
	if b == 'a' || b == 'e' || b == 'i' ||
		b == 'o' || b == 'u' ||
		b == 'A' || b == 'E' || b == 'I' ||
		b == 'O' || b == 'U' {
		return true
	}
	return false
}
