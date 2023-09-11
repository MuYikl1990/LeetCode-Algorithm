package main

import (
	"fmt"
	"sort"
)

func braceExpansion(s string) []string {
	var res []string
	var dfs1087 func(string, string)
	dfs1087 = func(s, tmp string) {
		if s == "" {
			res = append(res, tmp)
			return
		}
		if s[0] != '{' && s[0] != '}' {
			tmp += string(s[0])
			dfs1087(s[1:], tmp)
		} else {
			var char []string
			i := 1
			for ; i < len(s); i++ {
				if s[i] == '}' {
					break
				}
				if s[i] == ',' {
					continue
				}
				char = append(char, string(s[i]))
			}
			sort.Strings(char)
			for j := range char {
				tmp += char[j]
				dfs1087(s[i + 1:], tmp)
				tmp = tmp[:len(tmp) - 1]
			}
		}
	}

	dfs1087(s, "")
	return res
}

func main() {
	s := "{a,b,c}d{g,e,f}"
	res := braceExpansion(s)
	fmt.Println(res)
}
