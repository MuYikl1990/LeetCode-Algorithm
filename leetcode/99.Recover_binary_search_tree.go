package main

type TreeNode99 struct {
    Val int
    Left *TreeNode99
    Right *TreeNode99
}

func recoverTree(root *TreeNode99) {
	//var first, second, prevNode *TreeNode99
	//inorder99(root, &first, &second, &prevNode)
	//first.Val, second.Val = second.Val, first.Val

	var x, y, pre, post *TreeNode99
	for root != nil {
		if root.Left != nil {
			post = root.Left
			for post.Right != nil && post.Right != root {
				post = post.Right
			}

			if post.Right == nil {
				post.Right = root
				root = root.Left
			} else {
				if pre != nil && pre.Val > root.Val {
					y = root
					if x == nil {
						x = pre
					}
				}
				pre = root
				post.Right = nil
				root = root.Right
			}
		} else {
			if pre != nil && pre.Val > root.Val {
				y = root
				if x == nil {
					x = pre
				}
			}
			pre = root
			root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}

func inorder99(root *TreeNode99, first, second, prevNode **TreeNode99) {
	if root == nil { return }
	inorder99(root.Left, first, second, prevNode)
	if *prevNode != nil && (*prevNode).Val >= root.Val {
		if *first == nil {
			*first = *prevNode
		}
		*second = root
	}
	*prevNode = root
	inorder99(root.Right, first, second, prevNode)
}
