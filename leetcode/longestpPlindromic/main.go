package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababa"
	//s1 := "cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc"
	//s2 := "uwqrvqslistiezghcxaocjbhtktayupazvowjrgexqobeymperyxtfkchujjkeefmdngfabycqzlslocjqipkszmihaarekosdkwvsirzxpauzqgnftcuflzyqwftwdeizwjhloqwkhevfovqwyvwcrosexhflkcudycvuelvvqlbzxoajisqgwgzhioomucfmkmyaqufqggimzpvggdohgxheielsqucemxrkmmagozxhvxlwvtbbcegkvvdrgkqszgajebbobxnossfrafglxvryhvyfcibfkgpbsorqprfujfgbmbctsenvbzcvypcjubsnjrjvyznbswqawodghmigdwgijfytxbgpxreyevuprpztmjejkaqyhppchuuytkdsteroptkouuvmkvejfunmawyuezxvxlrjulzdikvhgxajohpzrshrnngesarimyopgqydcmsaciegqlpqnclpwcjqmhtmtwwtbkmtnntdllqbyyhfxsjyhugnjbebtxeljytoxvqvrxygmtogndrhlcmbmgiueliyfkkcuypvvzkomjrfhuhhnfbxeuvssvvllgukdolffukzwqaimxkngnjnmsbvwkajyxqntsqjkjqvwxnlxwjfiaofejtjcveqstqhdzgqistxwsgrovvwgorjaoosremqbzllgbgrwtqdggxnyvkivlcvnv"
	//s3 := "addb"
	s4 := "ac"
	fmt.Println(longestPalindrome(s4))
}
func longestPalindrome(s string) string {
	arr := convertStoArrray(s)
	n := len(arr)
	if n == 1 {
		return s
	}
	lenRel := []string{}
	len_max_string := 0
	index := 0
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if arr[i] == arr[j] {
				t := arr[i : j+1]
				str2 := strings.Join(t, "")
				if isPalindrome(str2) {
					lenRel = append(lenRel, str2)
					if len_max_string < len(str2) {
						len_max_string = len(str2)
						fmt.Println("max ", str2)
						index++
					} else {
						break
					}
				}
			}
		}
	}
	////t := 0
	////index := 0
	fmt.Println(lenRel)
	fmt.Println(index - 1)
	////for in, i := range lenRel {
	////	if len(i) > t {
	////		t = len(i)
	////		index = in
	////	}
	////}
	////if t != 0 {
	////	return lenRel[index]
	////}
	if index == 1 {
		return lenRel[index-1]
	} else {
		return lenRel[index]
	}
	return arr[0]
}
func convertStoArrray(s string) []string {
	chars := []rune(s)
	result := []string{}
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		result = append(result, char)
	}
	return result
}

//func checkDoiXung(s string) bool {
//	chars1 := []rune(s)
//	len1 := len(chars1)
//	for i := 0; i < len1/2; i++ {
//		if chars1[i] != chars1[len1-i-1] {
//			return false
//		}
//	}
//	return true
//}
func isPalindrome(s string) bool {
	var i, mid int
	n := len(s)
	if n%2 == 0 {
		mid = (n / 2) - 1
	} else {
		mid = n / 2
	}
	for i = 0; i <= mid; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}
