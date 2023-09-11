package main

type TreeNode958 struct {
    Val int
    Left *TreeNode958
    Right *TreeNode958
}

func isCompleteTree(root *TreeNode958) bool {
	if root == nil {
		return true
	}

	var queue []*TreeNode958
	prev := root
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		if prev == nil && node != nil {
			return false
		}
		if node != nil {
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
		prev = node
		queue = queue[1:]
	}
	return true
}
