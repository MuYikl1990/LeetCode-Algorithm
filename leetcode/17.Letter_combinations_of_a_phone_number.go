package main

import "fmt"

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	letterMap := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	var res []string
	var dfs func(string, int, []byte)
	dfs = func(digits string, index int, tmp []byte) {
		if index == len(digits) {
			res = append(res, string(tmp))
			return
		}
		value := letterMap[digits[index]]
		for j := 0; j < len(value); j++ {
			tmp = append(tmp, value[j])
			dfs(digits, index + 1, tmp)
			tmp = tmp[:len(tmp) - 1]
		}
	}
	dfs(digits, 0, []byte{})
	return res
}

func main() {
	digits := "22"
	res := letterCombinations(digits)
	fmt.Println(res)
}
