package main

type TreeNode145 struct {
    Val int
    Left *TreeNode145
    Right *TreeNode145
}

func postorderTraversal(root *TreeNode145) []int {
	var res []int
	var dfs func(*TreeNode145)
	dfs = func(root *TreeNode145) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		res = append(res, root.Val)
	}
	dfs(root)
	return res
}
