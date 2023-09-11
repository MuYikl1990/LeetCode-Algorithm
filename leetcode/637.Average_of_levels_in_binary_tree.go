package main

type TreeNode637 struct {
    Val int
    Left *TreeNode637
    Right *TreeNode637
}

func averageOfLevels(root *TreeNode637) []float64 {
	var res []float64
	var queue []*TreeNode637
	queue = append(queue, root)

	for len(queue) > 0 {
		length := len(queue)
		sum := 0
		for i := 0; i < length; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, float64(sum) / float64(length))
		queue = queue[length:]
	}
	return res
}
