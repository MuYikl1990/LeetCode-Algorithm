package main

type TreeNode226 struct {
    Val int
    Left *TreeNode226
    Right *TreeNode226
}

func invertTree(root *TreeNode226) *TreeNode226 {
	if root == nil {
		return root
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left
	return root
}
