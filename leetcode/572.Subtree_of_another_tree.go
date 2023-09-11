package main

type TreeNode572 struct {
    Val int
    Left *TreeNode572
    Right *TreeNode572
}

func isSubtree(root *TreeNode572, subRoot *TreeNode572) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
	return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(root, subRoot *TreeNode572) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
	return root.Val == subRoot.Val && isSameTree(root.Left, subRoot.Left) && isSameTree(root.Right, subRoot.Right)
}
