package main

import "fmt"

type TreeNode993 struct {
    Val int
    Left *TreeNode993
    Right *TreeNode993
}

func isCousins(root *TreeNode993, x int, y int) bool {
	var queue []*TreeNode993
	queue = append(queue, root)
	var xT, yT *TreeNode993

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			cur := queue[i]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
				if cur.Left.Val == x {
					xT = cur
				}
				if cur.Left.Val == y {
					yT = cur
				}
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
				if cur.Right.Val == x {
					xT = cur
				}
				if cur.Right.Val == y {
					yT = cur
				}
			}
		}
		if xT != nil && yT != nil && xT != yT {
			return true
		}
		if xT != nil || yT != nil {
			return false
		}
		queue = queue[n:]
	}
	return false
}

func main() {
	root, x, y := &TreeNode993{Val: 1, Left: &TreeNode993{Val: 2, Left: &TreeNode993{Val: 4}}, Right: &TreeNode993{Val: 3}}, 4, 3
	res := isCousins(root, x, y)
	fmt.Println(res)
}
