package main

type TreeNode1123 struct {
    Val int
    Left *TreeNode1123
    Right *TreeNode1123
}

func lcaDeepestLeaves(root *TreeNode1123) *TreeNode1123 {
	if root == nil {
	    return root
    }

    left := getDepth(root.Left)
    right := getDepth(root.Right)

    if left == right {
        return root
    } else if left > right {
        return lcaDeepestLeaves(root.Left)
    }
    return lcaDeepestLeaves(root.Right)
}

func getDepth(root *TreeNode1123) int {
    if root == nil {
        return 0
    }
    return max1123(getDepth(root.Left), getDepth(root.Right)) + 1
}

func max1123(a, b int) int {
    if a > b {
        return a
    }
    return b
}
