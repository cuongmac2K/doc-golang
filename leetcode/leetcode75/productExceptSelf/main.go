package main

import "fmt"

func main() {
	//a := []int{1, 2, 0, 3, 4}
	//a := []int{0, 0}
	a := []int{-1, 1, 0, 0, -3, 3}
	fmt.Println(productExceptSelf(a))
	fmt.Println(productExceptSelf3(a))
}
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := []int{}
	for i := 0; i < n; i++ {
		all := 1
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			all *= nums[j]
		}
		result = append(result, all)
	}
	return result
}
func productExceptSelf1(nums []int) []int {

	length := len(nums)
	answer := []int{}

	for i := 0; i < length; i++ {
		product := 1
		for j := 0; j < length; j++ {
			if i == j {
				continue
			}
			product *= nums[j]
		}
		answer = append(answer, product)
	}
	return answer
}
func productExceptSelf3(nums []int) []int {
	n := len(nums)
	result := []int{}
	sum := 1
	zero := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			zero++
			continue
		}
		sum *= nums[i]
	}
	for i := 0; i < n; i++ {
		if nums[i] == 0 && zero == 1 {
			result = append(result, sum)
			continue
		}
		if zero > 0 {
			result = append(result, 0)
			continue
		}
		result = append(result, sum/nums[i])
	}
	return result
}
