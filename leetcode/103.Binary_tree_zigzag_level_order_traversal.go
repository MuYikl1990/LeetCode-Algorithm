package main

import "fmt"

type TreeNode103 struct {
	Val int
	Left *TreeNode103
	Right *TreeNode103
}

func zigzagLevelOrder(root *TreeNode103) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	var queue []*TreeNode103
	depth := 0
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		var level []int
		for i := 0; i < n; i++ {
			cur := queue[i]
			level = append(level, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		if depth % 2 == 1 {
			for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
				level[j], level[k] = level[k], level[j]
			}
			res = append(res, level)
		} else {
			res = append(res, level)
		}
		depth++
		queue = queue[n:]
	}
	return res
}

func main() {
	root := &TreeNode103{Val: 3, Left: &TreeNode103{Val: 9}, Right: &TreeNode103{Val: 20, Left: &TreeNode103{Val: 15},
		Right: &TreeNode103{Val: 7}}}
	res := zigzagLevelOrder(root)
	fmt.Println(res)
}
