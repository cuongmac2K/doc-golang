package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(2))
}

func generateParenthesis(n int) (res []string) {
	cur := make([]byte, 2*n, 2*n)
	var l, r int
	var generator func()
	generator = func() {
		if l+r == 2*n {
			res = append(res, string(cur))
			return
		}
		if l < n {
			cur[l+r] = '('
			l++
			generator()
			l--
		}
		if r < l {
			cur[l+r] = ')'
			r++
			generator()
			r--
		}
	}
	generator()
	return
}
