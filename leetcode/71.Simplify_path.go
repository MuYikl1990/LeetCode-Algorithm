package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	files := strings.Split(path, "/")
	var stack []string
	for _, file := range files {
		if file == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack) - 1]
			}
		} else if file != "." && file != "" {
			stack = append(stack, file)
		}
	}
	res := "/" + strings.Join(stack, "/")
	return res
}

func main() {
	path := "/home//foo/"
	res := simplifyPath(path)
	fmt.Println(res)
}
