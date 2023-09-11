package main

type TreeNode144 struct {
    Val int
    Left *TreeNode144
    Right *TreeNode144
}

func preorderTraversal(root *TreeNode144) []int {
	var res []int
	dfs144(root, &res)
	return res
}

func dfs144(root *TreeNode144, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	dfs144(root.Left, res)
	dfs144(root.Right, res)
}
