package main

import "fmt"

type TreeNode199 struct {
    Val int
    Left *TreeNode199
    Right *TreeNode199
}

func rightSideView(root *TreeNode199) []int {
	var queue []*TreeNode199
	var res []int
	if root == nil {
		return res
	}
	queue = append(queue, root)

	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if i == 0 {
				res = append(res, queue[i].Val)
			}
		}
		queue = queue[length:]
	}
	return res
}

func main() {
	root := &TreeNode199{Val: 1, Left: &TreeNode199{Val: 2, Right: &TreeNode199{Val: 5}}, Right: &TreeNode199{Val: 3,
			Right: &TreeNode199{Val: 4}}}
	res := rightSideView(root)
	fmt.Println(res)
}
