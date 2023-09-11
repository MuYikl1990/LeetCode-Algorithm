package main

import "fmt"

type TreeNode543 struct {
	Val int
	Left *TreeNode543
	Right *TreeNode543
}

func diameterOfBinaryTree(root *TreeNode543) int {
	var res int
	dfs543(root, &res)
	return res
}

func dfs543(root *TreeNode543, res *int) int {
	if root == nil {
		return 0
	}

	left := dfs543(root.Left, res)
	right := dfs543(root.Right, res)
	*res = max543(*res, left + right)
	return max543(left, right) + 1
}

func max543(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode543{Val: 1, Left: &TreeNode543{Val: 2, Left: &TreeNode543{Val: 4}, Right: &TreeNode543{Val: 5}},
		Right: &TreeNode543{Val: 3}}
	res := diameterOfBinaryTree(root)
	fmt.Println(res)
}
