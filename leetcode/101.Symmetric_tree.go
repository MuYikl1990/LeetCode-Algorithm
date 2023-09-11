package main

type TreeNode101 struct {
    Val int
    Left *TreeNode101
    Right *TreeNode101
}

func isSymmetric(root *TreeNode101) bool {
	if root == nil {
		return true
	}

	return dfs101(root.Left, root.Right)
}

func dfs101(left, right *TreeNode101) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return dfs101(left.Left, right.Right) && dfs101(left.Right, right.Left)
}
