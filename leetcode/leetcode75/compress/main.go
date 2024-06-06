package main

import (
	"strconv"
)

func main() {
	//s := []string{"a", "a", "b", "b", "c", "c", "c"}

	//fmt.Println(compress())
}
func compress(chars []byte) int {
	var res int
	for i := 0; i < len(chars); {
		j := i + 1
		for ; j < len(chars); j++ {
			if chars[j] != chars[j-1] {
				break
			}
		}
		chars[res] = chars[j-1]
		res++
		n := j - i
		if n > 1 {
			num := strconv.Itoa(n)
			for k := range num {
				chars[res] = num[k]
				res++
			}
		}
		i = j
	}
	return res
}
