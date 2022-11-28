package main

import "fmt"

func main() {
	s := "{[]}"
	fmt.Println(isValid(s))
}
func isValid(s string) bool {
	char := []rune(s)
	ngoactruoc := ""
	ngoacsau := ""
	n := len(char)
	flag1 := true
	flag2 := true
	if n%2 == 0 {
		for i := 0; i < n-1; i++ {
			if i%2 == 0 {
				ngoactruoc = string(char[i])
				ngoacsau = string(char[i+1])
				//fmt.Println("++", ngoactruoc, ngoacsau, i+1)
				if daoNguocMo(ngoactruoc) != ngoacsau {
					flag1 = false
				}
			}
		}
	} else {
		flag1 = false
	}
	//	fmt.Println(n)
	if flag1 == false {
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				ngoactruoc = string(char[i])
				ngoacsau = string(char[n-i-1])
				fmt.Println("++", ngoactruoc, daoNguocMo(ngoacsau), i+1)
				if ngoactruoc != daoNguocMo(ngoacsau) {
					fmt.Println("false", ngoactruoc, ngoacsau)
					flag2 = false
				}
			}
		}
	}
	if flag1 == false && flag2 == false {
		return false
	}
	return true
}

func daoNguocMo(s string) string {
	switch s {
	case "(":
		return ")"
	case "[":
		return "]"
	case "{":
		return "}"
	}
	return s
}
