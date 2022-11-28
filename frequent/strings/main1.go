package main

import "fmt"

func main() {
	s := "Thị Trấn Tây Đằng, Huyện Ba Vì, Thành phố Hà Nội"
	s1 := "Thị trấn Tây Đằng, Huyện Ba Vì, Thành phố Hà Nội"
	if s == s1 {
		fmt.Println(1)
	} else {
		fmt.Println(2)
	}
}
