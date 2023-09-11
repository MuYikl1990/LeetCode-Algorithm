package main

type TreeNode700 struct {
    Val int
    Left *TreeNode700
    Right *TreeNode700
}

func searchBST(root *TreeNode700, val int) *TreeNode700 {
    var res *TreeNode700
    var dfs func(*TreeNode700)
    dfs = func(root *TreeNode700) {
        if root == nil {
            return
        }
        dfs(root.Left)
        if root.Val == val {
            res = root
        }
        dfs(root.Right)
        return
    }
    dfs(root)
    return res
}

func searchBST1(root *TreeNode700, val int) *TreeNode700 {
    if root == nil || root.Val == val {
        return root
    }

    left := searchBST(root.Left, val)
    right := searchBST(root.Right, val)

    if left != nil {
        return left
    } else {
        return right
    }
}
