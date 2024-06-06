package main

func main() {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
}
func moveZeroes(nums []int) {
	index := 0

	for _, n := range nums {
		if n != 0 {
			nums[index] = n
			index++
		}
	}
	for index < len(nums) {
		nums[index] = 0
		index++
	}
}
