package main

import (
	"fmt"
	"reflect"
)

func generateParenthesis(n int) []string {
	var res []string
	dfs22(n, 0, 0, &res, "")
	return res
}

func dfs22(n, left, right int, res *[]string, tmp string) {
	if left == n && right == n {
		*res = append(*res, tmp)
		return
	}

	if left < right {
		return
	}

	if left < n {
		dfs22(n, left + 1, right, res, tmp + "(")
	}

	if right < n {
		dfs22(n, left, right + 1, res, tmp + ")")
	}
	return
}

func main() {
	n := 3
	res := generateParenthesis(n)
	target := []string{"((()))","(()())","(())()","()(())","()()()"}
	fmt.Println(res, reflect.DeepEqual(res, target))
}
