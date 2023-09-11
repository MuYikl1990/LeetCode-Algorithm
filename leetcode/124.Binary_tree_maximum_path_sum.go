package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	sum := math.MinInt
	dfs(root, &sum)
	return sum
}

func dfs(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left, sum)
	right := dfs(root.Right, sum)
	lrpath := root.Val + Max(0, left) + Max(0, right)
	ret := root.Val + Max(0, Max(left, right))
	*sum = Max(*sum, Max(lrpath, ret))
	return ret
}

func main() {
	root := &TreeNode{Val: -10, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15},
		Right: &TreeNode{Val: 7}}}
	res := maxPathSum(root)
	fmt.Println(res)
}
