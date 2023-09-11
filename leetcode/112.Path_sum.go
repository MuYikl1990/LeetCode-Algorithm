package main

type TreeNode112 struct {
    Val int
    Left *TreeNode112
    Right *TreeNode112
}

func hasPathSum(root *TreeNode112, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}

	return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
}
