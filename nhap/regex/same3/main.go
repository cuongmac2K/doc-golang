package main

import (
	"fmt"
	"regexp"
)

func main() {
	//pattern := `.*namespaces/(.*?)/unzip-files`
	pattern := `.*namespaces/(.*?)/get-capacity`
	re := regexp.MustCompile(pattern)

	testStrings := []string{
		"abc/namespaces/test/get-capacity",
		"xyz/namespaces/123/get-capacity",
		"prefix/namespaces/test/get-capacity/extra",
		"other/namespaces//get-capacity",
		"foo/namespaces/test/other-capacity",
		"namespaces/test/get-capacity",
	}
	fmt.Printf("\n")
	for _, str := range testStrings {
		if re.MatchString(str) {
			fmt.Printf("Matched: %s\n", str)
		} else {
			fmt.Printf("Did not match: %s\n", str)
		}
	}
}
