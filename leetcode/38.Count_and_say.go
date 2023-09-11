package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		start, end := 0, 0
		tmp := ""
		for end < len(res) {
			for end < len(res) && res[start] == res[end] {
				end++
			}
			tmp += strconv.Itoa(end - start) + string(res[start])
			start = end
		}
		res = tmp
	}
	return res
}

func main() {
	n := 5
	res := countAndSay(n)
	fmt.Println(res)
}
