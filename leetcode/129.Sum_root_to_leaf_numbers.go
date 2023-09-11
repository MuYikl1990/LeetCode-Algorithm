package main

type TreeNode129 struct {
    Val int
    Left *TreeNode129
    Right *TreeNode129
}

func sumNumbers(root *TreeNode129) int {
    return sum(root, 0)
}

func sum(root *TreeNode129, num int) int {
    if root == nil {
        return 0
    }

    num = 10 * num + root.Val
    if root.Left == nil && root.Right == nil {
        return num
    }
    return sum(root.Left, num) + sum(root.Right, num)
}
