package main

import "fmt"

type TreeNode94 struct {
	Val int
	Left *TreeNode94
	Right *TreeNode94
}

func inorderTraversal(root *TreeNode94) []int {
	var res []int
	inorder(root, &res)
	return res
}

func inorder(root *TreeNode94, res *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}

func main() {
	root := &TreeNode94{Val: 2, Left: &TreeNode94{Val: 1}, Right: &TreeNode94{Val: 4, Left: &TreeNode94{Val: 3},
		Right: &TreeNode94{Val: 6}}}
	res := inorderTraversal(root)
	fmt.Println(res)
}
