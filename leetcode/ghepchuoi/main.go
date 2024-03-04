package main

import "fmt"

func main() {
	//word1 := "ace"
	//word2 := "bdfghiklm"
	word1 := "acegh"
	word2 := "bdf"
	fmt.Println(ghep(word1, word2))
}
func ghep(w1, w2 string) (s string) {
	if len(w2) > len(w1) {
		n := len(w1)
		for i := 0; i < n; i++ {
			s = s + string(w1[i]) + string(w2[i])
		}
		s = s + w2[n:]
	} else {
		n := len(w2)
		for i := 0; i < n; i++ {
			s = s + string(w1[i]) + string(w2[i])
		}
		s = s + w1[n:]
	}
	return s
}
