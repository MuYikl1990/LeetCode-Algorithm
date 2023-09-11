package main

import (
	"fmt"
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	var res []string

	for i := 0; i < len(folder); {
		fd := folder[i]
		res = append(res, fd)
		compare := fd + "/"
		i++
		for i < len(folder) && strings.Index(folder[i], compare) == 0 {
			i++
		}
	}
	return res
}

func main() {
	folder := []string{"/a","/a/b","/c/d","/c/d/e","/c/f"}
	res := removeSubfolders(folder)
	fmt.Println(res)
}
