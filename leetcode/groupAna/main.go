package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(str))
}
func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	//var arr [3][4]string
	arr := make([][]string, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]string, n)
	}
	fmt.Println(arr)
	for i := 0; i < n; i++ {
		index := 0
		arr[i][index] = strs[i]
		for j := 1; j < n; j++ {
			if checkExits(strs[i], strs[j]) {
				arr[i][index] = strs[j]
				strs = strs[:j+copy(strs[j:], strs[j+1:])]
				index++
				n--
			}
		}
	}
	return arr
}
func checkExits(s1, s2 string) bool {
	ar1 := strings.Split(s1, "")
	ar2 := strings.Split(s2, "")
	sort.Strings(ar1)
	sort.Strings(ar2)
	justString1 := fmt.Sprint(ar1)
	justString2 := fmt.Sprint(ar2)
	if justString1 != justString2 {
		return false
	}
	return true
}
