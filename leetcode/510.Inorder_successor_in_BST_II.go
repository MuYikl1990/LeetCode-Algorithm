package main

type TreeNode510 struct {
	Val int
	Left *TreeNode510
	Right *TreeNode510
	Parent *TreeNode510
}

func inorderSuccessor(node *TreeNode510) *TreeNode510 {
	if node == nil {
		return nil
	}
	if node.Right != nil {
		node = node.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
	}
	for node.Parent != nil && node.Parent.Right == node {
		node = node.Parent
	}
	return node.Parent
}
