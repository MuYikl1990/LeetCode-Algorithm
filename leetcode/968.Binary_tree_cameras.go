package main

type TreeNode968 struct {
    Val int
    Left *TreeNode968
    Right *TreeNode968
}

func minCameraCover(root *TreeNode968) int {
	res := 0
	if dfs968(root, &res) == 0 {
		res++
	}
	return res
}

func dfs968(root *TreeNode968, res *int) int {
	if root == nil {
		return 2
	}

	left := dfs968(root.Left, res)
	right := dfs968(root.Right, res)

	if left == 2 && right == 2 {
		return 0
	}

	if left == 0 || right == 0 {
		*res++
		return 1
	}

	if left == 1 || right == 1 {
		return 2
	}

	return 2
}
