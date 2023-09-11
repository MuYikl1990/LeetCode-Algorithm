package main

import (
	"fmt"
	"sort"
	"strings"
)

func braceExpansionII(expression string) []string {
	strs := make(map[string]struct{})
	var dfs1096 func(string)
	dfs1096 = func(expr string) {
		right := strings.Index(expr, "}")
		if right == -1 {
			strs[expr] = struct{}{}
			return
		}
		left := strings.LastIndex(expr[:right], "{")
		pre, post := expr[:left], expr[right + 1:]
		for _, s := range strings.Split(expr[left + 1:right], ",") {
			dfs1096(pre + s + post)
		}
	}
	dfs1096(expression)
	var res []string
	for str := range strs {
		res = append(res, str)
	}
	sort.Strings(res)
	return res
}

func main() {
	expression := "{{a,z},a{b,c},{ab,z}}"
	res := braceExpansionII(expression)
	fmt.Println(res)
}
