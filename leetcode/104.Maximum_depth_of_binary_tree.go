package main

type TreeNode104 struct {
    Val int
    Left *TreeNode104
    Right *TreeNode104
}

func maxDepth(root *TreeNode104) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return max104(left, right) + 1
}

func max104(a, b int) int {
	if a > b {
		return a
	}
	return b
}
