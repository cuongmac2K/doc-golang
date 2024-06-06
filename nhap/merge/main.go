package main

import "slices"

func main() {

}

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	for i := 1; i < len(intervals); i++ {
		prev := intervals[i-1]
		cur := intervals[i]

		if prev[1] >= cur[0] {
			intervals[i] = []int{prev[0], max(prev[1], cur[1])}
			intervals = append(intervals[:i-1], intervals[i:]...)
			i--
		}
	}

	return intervals
}
