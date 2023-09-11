package main

type TreeNode513 struct {
    Val int
    Left *TreeNode513
    Right *TreeNode513
}

func findBottomLeftValue(root *TreeNode513) int {
	res := 0
	var queue []*TreeNode513
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if i == 0 {
				res = queue[i].Val
			}
		}
		queue = queue[n:]
	}
	return res
}
