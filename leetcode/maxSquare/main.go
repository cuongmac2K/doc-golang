package main

import "fmt"

func main() {
	arr := [][]byte{{1, 0, 1, 0, 0}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}}
	fmt.Println(maximalSquare(arr))
}
func maximalSquare(matrix [][]byte) int {
	n := len(matrix)
	if n == 1 {
		for i := 0; i < n; i++ {
			if matrix[1][i] == 1 {
				return 1
			}
		}
	}
	maxWide := 0
	MaxHeight := 0
	dai := 0
	rong := 0
	m := len(matrix[0])
	for i := 0; i < n-1; i++ {
		fmt.Println(matrix[i])
		for j := 0; j < m-1; j++ {

			if matrix[i][j] == 1 && matrix[i][j+1] == 1 {
				dai++
				if dai > maxWide {
					maxWide = dai
				}
			} else {
				dai = 0
			}

			if matrix[i+1][j] == 1 && matrix[i][j] == 1 {
				rong++
				if rong > MaxHeight {
					MaxHeight = rong
				}
			} else {
				rong = 0
			}
		}
	}
	fmt.Println(maxWide)
	fmt.Println(MaxHeight)
	return maxWide * MaxHeight
}
