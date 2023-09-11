package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var res []string
	dfs93(s, &res, []string{}, 0)
	return res
}

func dfs93(s string, res *[]string, list []string, start int) {
	if len(list) == 4 && start == len(s) {
		tmp := strings.Join(list, ".")
		*res = append(*res, tmp)
		return
	}

	if len(list) == 4 && start < len(s) {
		return
	}

	for length := 1; length <= 3; length++ {
		if start + length - 1 >= len(s) {
			return
		}

		if length != 1 && s[start] == '0' {
			return
		}

		if n, _ := strconv.Atoi(s[start:start+length]); n > 255 {
			return
		}
		list = append(list, s[start:start+length])
		dfs93(s, res, list, start+length)
		list = list[:len(list)-1]
	}
	return
}

func main() {
	s := "101023"
	res := restoreIpAddresses(s)
	fmt.Println(res)
}
