package main

type TreeNode701 struct {
    Val int
    Left *TreeNode701
    Right *TreeNode701
}

func insertIntoBST(root *TreeNode701, val int) *TreeNode701 {
	node := &TreeNode701{Val: val}
	if root == nil {
		return node
	}

	cur := root
	for {
		if cur.Val > val {
			if cur.Left == nil {
				cur.Left = node
				break
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = node
				break
			}
			cur = cur.Right
		}
	}
	return root
}
