package main

import "fmt"

func main() {
	s := "VX"
	//fmt.Println(romanToInt(s))
	fmt.Println(romanToInt1(s))

}
func romanToInt(s string) int {
	translation := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	res := 0
	i := 0
	for i < len(s) {
		offset := 1
		for {
			if i+offset > len(s) {
				break
			}
			_, ok := translation[s[i:i+offset]]
			if !ok {
				break
			}
			offset++
		}
		res += translation[s[i:i+offset-1]]
		i += offset - 1
	}
	return res
}
func romanToInt1(s string) int {
	romans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romans[string(s[i])] < romans[string(s[i+1])] {
			result -= romans[string(s[i])]

		} else {
			result += romans[string(s[i])]
		}
	}
	return result
}
