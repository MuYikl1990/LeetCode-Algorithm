package main

import (
	"fmt"
)

type TreeNode111 struct {
    Val int
    Left *TreeNode111
    Right *TreeNode111
}

func minDepth(root *TreeNode111) int {
	//if root == nil {
	//	return 0
	//}
	//
	//res := math.MaxInt
	//var dfs func(int, *TreeNode111)
	//dfs = func(depth int, root *TreeNode111) {
	//	if root.Left == nil && root.Right == nil {
	//		res = min111(res, depth)
	//		return
	//	}
	//
	//	dfs(depth + 1, root.Left)
	//	dfs(depth + 1, root.Right)
	//	return
	//}
	//dfs(1, root)
	//return res

	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}
	res := min111(left, right) + 1
	return res
}

func min111(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode111{Val: 3, Left: &TreeNode111{Val: 9}, Right: &TreeNode111{Val: 20, Left: &TreeNode111{Val: 15}, Right: &TreeNode111{Val: 7}}}
	res := minDepth(root)
	fmt.Println(res)
}
