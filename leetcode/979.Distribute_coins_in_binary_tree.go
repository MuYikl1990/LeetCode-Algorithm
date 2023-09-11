package main

import "fmt"

type TreeNode979 struct {
    Val int
    Left *TreeNode979
    Right *TreeNode979
}

func distributeCoins(root *TreeNode979) int {
	var res int
	postOrder(root, &res)
	return res
}

func postOrder(root *TreeNode979, res *int) int {
	if root == nil {
		return 0
	}
	left := postOrder(root.Left, res)
	right := postOrder(root.Right, res)
	*res += abs979(left) + abs979(right)
	return root.Val + left + right - 1
}

func abs979(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	root := &TreeNode979{Val: 0, Left: &TreeNode979{Val: 3}, Right: &TreeNode979{Val: 0}}
	res := distributeCoins(root)
	fmt.Println(res)
}
