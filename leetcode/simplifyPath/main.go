package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/home//foo/"
	fmt.Println(simplifyPath(path))
}
func simplifyPath(path string) string {
	fields := strings.Split(strings.Trim(path, "/"), "/")
	var ans []string

	for _, f := range fields {
		switch f {
		case "..":
			if len(ans) > 0 {
				ans = ans[:len(ans)-1]
			}
			break
		case ".":
		case "":
			break
		default:
			ans = append(ans, f)
		}
	}

	return fmt.Sprintf("/%s", strings.Join(ans, "/"))

}
