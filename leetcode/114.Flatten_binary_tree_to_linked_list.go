package main

type TreeNode114 struct {
    Val int
    Left *TreeNode114
    Right *TreeNode114
}

func flatten(root *TreeNode114) {
    var pre *TreeNode114
    var dfs func(*TreeNode114)
    dfs = func(root *TreeNode114) {
        if root == nil {
            return
        }
        dfs(root.Right)
        dfs(root.Left)
        root.Right = pre
        root.Left = nil
        pre = root
    }
    dfs(root)
}
