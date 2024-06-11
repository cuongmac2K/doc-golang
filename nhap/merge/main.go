package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Compile the regex with a negative lookahead assertion
	regex, err := regexp.CompilePOSIX(`^(?!\/\.)[A-Za-z0-9_.@]{1,255}$`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	testStrings := []string{
		"validString_1@example.com",
		"/.invalidStart",
		"another.valid_string@domain.com",
		"invalid/string",
		"valid_string@example.com",
	}

	for _, str := range testStrings {
		if regex.MatchString(str) {
			fmt.Printf("Matched: %s\n", str)
		} else {
			fmt.Printf("Did not match: %s\n", str)
		}
	}
}
