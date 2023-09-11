package main

type Node426 struct {
	Val int
	Left *Node426
	Right *Node426
}

func treeToDoublyList(root *Node426) *Node426 {
	var pre, head *Node426
	dfs426(root, pre, head)
	head.Left, pre.Right = pre, head
	return head
}

func dfs426(root, pre, head *Node426) {
	if root == nil {
		return
	}
	dfs426(root.Left, pre, head)
	if pre != nil {
		pre.Right = root
	} else {
		head = root
	}
	root.Left = pre
	pre = root
	dfs426(root.Right, pre, head)
}
