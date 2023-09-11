package main

import (
	"fmt"
	"sort"
)

func permutation(s string) []string {
	char := []byte(s)
	sort.Slice(char, func(i, j int) bool {
		return char[i] <= char[j]
	})
	used := make([]bool, len(char))
	var res []string
	var backtrace38 func([]byte, string)
	backtrace38 = func(char []byte, path string) {
		if len(path) == len(char) {
			res = append(res, path)
			return
		}

		for i := 0; i < len(char); i++ {
			if i > 0 && char[i] == char[i - 1] && !used[i - 1] {
				continue
			}
			if !used[i] {
				used[i] = true
				backtrace38(char, path + string(char[i]))
				used[i] = false
			}
		}
	}
	backtrace38(char, "")
	return res
}

func main() {
	s := "suvyls"
	res := permutation(s)
	fmt.Println(res)
}
