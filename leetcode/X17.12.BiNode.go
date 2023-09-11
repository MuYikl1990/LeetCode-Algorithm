package main

import "fmt"

type TreeNode12 struct {
    Val int
    Left *TreeNode12
    Right *TreeNode12
}

func convertBiNode(root *TreeNode12) *TreeNode12 {
	dummy := &TreeNode12{}
	var preNode *TreeNode12
	var dfs12 func(*TreeNode12)
	dfs12 = func(root *TreeNode12) {
		if root == nil {
			return
		}
		dfs12(root.Left)
		if preNode == nil {
			preNode = root
			dummy.Right = root
		} else {
			preNode.Right = root
			preNode = root
		}
		root.Left = nil
		dfs12(root.Right)
		return
	}

	dfs12(root)
	return dummy.Right
}

func convertBiNode1(root *TreeNode12) *TreeNode12 {
	var preNode *TreeNode12
	var dfs12 func(*TreeNode12)
	dfs12 = func(root *TreeNode12) {
		if root == nil {
			return
		}
		dfs12(root.Right)
		root.Right = preNode
		preNode = root
		dfs12(root.Left)
		root.Left = nil
		return
	}

	dfs12(root)
	return preNode
}

func main() {
	root := &TreeNode12{Val: 4, Left: &TreeNode12{Val: 2, Left: &TreeNode12{Val: 1, Left: &TreeNode12{Val: 0}}, Right: &TreeNode12{Val: 3}}, Right: &TreeNode12{Val: 5, Right: &TreeNode12{Val: 6}}}
	res := convertBiNode(root)
	fmt.Println(res.Right.Val)
}
