package main

import (
	"fmt"
)

type TreeNode988 struct {
    Val int
    Left *TreeNode988
    Right *TreeNode988
}

func smallestFromLeaf(root *TreeNode988) string {
	var res string
	dfs988(root, "", &res)
	return res
}

func dfs988(root *TreeNode988, tmp string, res *string) {
	if root == nil {
		return
	}

	tmp = string(rune('a' + root.Val)) + tmp
	if root.Left == nil && root.Right == nil {
		if *res > tmp || *res == "" {
			*res = tmp
		}
		return
	}
	dfs988(root.Left, tmp, res)
	dfs988(root.Right, tmp, res)
	return
}

func main() {
	root := &TreeNode988{Val: 0, Left: &TreeNode988{Val: 1, Left: &TreeNode988{Val: 3}, Right: &TreeNode988{Val: 4}},
		Right: &TreeNode988{Val: 2, Left: &TreeNode988{Val: 3}, Right: &TreeNode988{Val: 4}}}
	res := smallestFromLeaf(root)
	fmt.Println(res)
}
