package main

import "fmt"

func main() {
	//flo := []int{1, 0, 0, 0, 0, 0, 0, 0, 1}
	//n := 4
	flo := []int{1, 0, 0, 0, 0, 1}
	n := 2
	fmt.Println(canPlaceFlowers(flo, n))
}

//	func canPlaceFlowers(flowerbed []int, n int) bool {
//		nf := len(flowerbed)
//		if nf < 4 {
//			return false
//		}
//		maxFlower := 0
//
//		for i := 0; i < nf-2; i++ {
//			if flowerbed[i] == flowerbed[i+1] && flowerbed[i+1] == flowerbed[i+2] && flowerbed[i+2] == 0 {
//				maxFlower++
//				i = i + 1
//			}
//		}
//		if maxFlower >= n {
//			return true
//		}
//		return false
//	}
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i, plot := range flowerbed {
		if plot == 1 {
			continue
		}

		if (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}

	return n <= 0
}
