package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 4
	nums2 := []int{2, 5, 6}
	n := 3
	fmt.Println(merge(nums1, m, nums2, n))
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	slice := nums1[:m]
	slice = append(slice, nums2[:n]...)
	sort.Ints(slice)
	return slice
}
