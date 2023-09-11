package main

import "fmt"

type TreeNode54 struct {
    Val int
    Left *TreeNode54
    Right *TreeNode54
}

func kthLargest(root *TreeNode54, k int) int {
	var res int
	inorder54(root, &k, &res)
	return res
}

func inorder54(root *TreeNode54, k *int, res *int) {
	if root == nil {
		return
	}
	inorder54(root.Right, k, res)
	*k--
	if *k == 0 {
		*res = root.Val
		return
	}
	inorder54(root.Left, k, res)
}

func main() {
	root, k := &TreeNode54{Val: 3, Left: &TreeNode54{Val: 1, Right: &TreeNode54{Val: 2}}, Right: &TreeNode54{Val: 4}}, 2
	res := kthLargest(root, k)
	fmt.Println(res)
}
