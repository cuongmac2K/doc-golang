package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
}
func rotate(matrix [][]int) {
	n := len(matrix)

	// Xoay ma trận theo từng lớp hoán vị
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer

		// Xoay từng phần tử trong lớp hoán vị
		for i := first; i < last; i++ {
			offset := i - first

			// Lưu giữ giá trị của phần tử trên cạnh trên
			top := matrix[first][i]

			// Gán giá trị của phần tử trên cạnh trái cho cạnh trên
			matrix[first][i] = matrix[last-offset][first]

			// Gán giá trị của phần tử trên cạnh dưới cho cạnh trái
			matrix[last-offset][first] = matrix[last][last-offset]

			// Gán giá trị của phần tử trên cạnh phải cho cạnh dưới
			matrix[last][last-offset] = matrix[i][last]

			// Gán giá trị của phần tử trên cạnh trên cho cạnh phải
			matrix[i][last] = top
		}
	}

	fmt.Println(matrix)
}
