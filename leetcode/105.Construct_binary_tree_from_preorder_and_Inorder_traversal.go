package main

type TreeNode105 struct {
    Val int
    Left *TreeNode105
    Right *TreeNode105
}

func buildTree105(preorder []int, inorder []int) *TreeNode105 {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode105{Val: preorder[0]}
	mid := 0
	for i, val := range inorder {
		if val == root.Val {
			mid = i
		}
	}
	root.Left = buildTree105(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree105(preorder[mid+1:], inorder[mid+1:])
	return root
}
