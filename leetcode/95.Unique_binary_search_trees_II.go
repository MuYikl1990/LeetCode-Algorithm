package main

type TreeNode95 struct {
    Val int
    Left *TreeNode95
    Right *TreeNode95
}

func generateTrees(n int) []*TreeNode95 {
	if n == 0 {
		return []*TreeNode95{}
	}
	return makeTree(1, n)
}

func makeTree(start, end int) []*TreeNode95 {
	var res []*TreeNode95
	if start > end {
		res = append(res, nil)
		return res
	}

	if start == end {
		res = append(res, &TreeNode95{Val: start})
		return res
	}

	for i := start; i <= end; i++ {
		left := makeTree(start, i - 1)
		right := makeTree(i + 1, end)
		for _, l := range left {
			for _, r := range right {
				node := &TreeNode95{Val: i}
				node.Left = l
				node.Right = r
				res = append(res, node)
			}
		}
	}
	return res
}
