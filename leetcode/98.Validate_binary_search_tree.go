package main

import (
	"fmt"
	"math"
)

type TreeNode98 struct {
	Val int
	Left *TreeNode98
	Right *TreeNode98
}

func isValidBST(root *TreeNode98) bool {
	pre := math.MinInt
	return dfs98(root, &pre)
}

func dfs98(root *TreeNode98, pre *int) bool {
	if root == nil {
		return true
	}

	if !dfs98(root.Left, pre) {
		return false
	}

	if root.Val <= *pre {
		return false
	}
	*pre = root.Val

	return dfs98(root.Right, pre)
}

func main() {
	root := &TreeNode98{Val: 5, Left: &TreeNode98{Val: 1}, Right: &TreeNode98{Val: 4, Left: &TreeNode98{Val: 3},
		Right: &TreeNode98{Val: 6}}}
	res := isValidBST(root)
	fmt.Println(res)
}
