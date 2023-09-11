package main

type TreeNode100 struct {
    Val int
    Left *TreeNode100
    Right *TreeNode100
}

func isSameTree100(p *TreeNode100, q *TreeNode100) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree100(p.Left, q.Left) && isSameTree100(p.Right, q.Right)
}
