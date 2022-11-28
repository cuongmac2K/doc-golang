package main

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

func main() {

	text := "$$"

	fmt.Println(ValidName(text))
}
func ValidName(e string) bool {
	emailRegex := regexp.MustCompile(`[0-9a-zA-Z_.-]`)
	return emailRegex.MatchString(e)
}

//[a-zA-Z_0-9_.-]
//func ValidateName(name string) bool {
//	str := strings.Split(name, "")
//	for _, s := range str {
//		switch s {
//		case "!":
//			return false
//		case "@":
//			return false
//		case "#":
//			return false
//		case "$":
//			return false
//		case "%":
//			return false
//		case "":
//			return false
//		case "&":
//			return false
//		case "*":
//			return false
//		case "(":
//			return false
//		case ")":
//			return false
//		case "_":
//			return false
//		case "`":
//			return false
//		case "~":
//			return false
//		case ".":
//			return false
//		case ",":
//			return false
//		case "[":
//			return false
//		case "?":
//			return false
//		case "?":
//			return false
//
//		}
//
//	}
//	return true
//}
func vali(text string) bool {
	kitu := []string{"~", "@", "!", "#", "$", "%", "^", "&", "*", "(", ")", "_", "-", "=", "+", "[", "]", "{", "}", ";", ":", "|", ".", "`", "?", "<", ">", ","}
	var wg sync.WaitGroup
	tg := make(chan string)
	for _, i := range kitu {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i string) {
			if strings.Contains(text, i) {
				tg <- "false"
			} else {
				tg <- "true"
			}
		}(&wg, i)
	}
	if <-tg == "false" {
		return false
	}
	if <-tg == "true" {
		return true
	}
	wg.Wait()
	return true
}
