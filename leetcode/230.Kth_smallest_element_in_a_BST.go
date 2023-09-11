package main

import "fmt"

type TreeNode230 struct {
    Val int
    Left *TreeNode230
    Right *TreeNode230
}

func kthSmallest(root *TreeNode230, k int) int {
	var stack []*TreeNode230
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		stack, root = stack[:len(stack) - 1], stack[len(stack) - 1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

func main() {
	root := &TreeNode230{Val: 3, Left: &TreeNode230{Val: 1, Left: nil, Right: &TreeNode230{Val: 2}}, Right: &TreeNode230{Val: 4}}
	res := kthSmallest(root, 2)
	fmt.Println(res)
}