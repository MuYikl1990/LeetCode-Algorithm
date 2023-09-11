package main

type TreeNode617 struct {
    Val int
    Left *TreeNode617
    Right *TreeNode617
}

func mergeTrees(root1 *TreeNode617, root2 *TreeNode617) *TreeNode617 {
    if root1 == nil {
        return root2
    }
    if root2 == nil {
        return root1
    }

    root1.Val += root2.Val
    root1.Left = mergeTrees(root1.Left, root2.Left)
    root1.Right = mergeTrees(root1.Right, root2.Right)
    return root1
}
