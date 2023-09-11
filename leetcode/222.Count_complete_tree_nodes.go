package main

type TreeNode222 struct {
    Val int
    Left *TreeNode222
    Right *TreeNode222
}

func countNodes(root *TreeNode222) int {
	left, right := root, root
	hl, hr := 0, 0

	for left != nil {
		left = left.Left
		hl++
	}

	for right != nil {
		right = right.Right
		hr++
	}

	if hl == hr {
		return 1 << hl - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
