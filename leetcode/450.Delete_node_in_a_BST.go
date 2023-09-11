package main

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			pre := root
			node := root.Right
			for node.Left != nil {
				pre = node
				node = node.Left
			}
			root.Val = node.Val
			if pre.Left.Val == node.Val {
				pre.Left = node.Right
			} else {
				pre.Right = node.Right
			}
		}
	}
	return root
}
