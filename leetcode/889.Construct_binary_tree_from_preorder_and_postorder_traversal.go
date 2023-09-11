package main

type TreeNode889 struct {
    Val int
    Left *TreeNode889
    Right *TreeNode889
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode889 {
    n := len(preorder)
    var dfs func([]int, int, int, []int, int, int) *TreeNode889
    dfs = func(preorder []int, preStart, preEnd int, postorder []int, postStart, postEnd int) *TreeNode889 {
        if preStart > preEnd || postStart > postEnd {
            return nil
        }
        root := &TreeNode889{Val: preorder[preStart]}
        if preStart == preEnd {
            return root
        }
        index := postStart
        for i := postStart; i <= postEnd; i++ {
            if postorder[i] == preorder[preStart + 1] {
                index = i
                break
            }
        }
        root.Left = dfs(preorder, preStart + 1, preStart + 1 + index - postStart, postorder, postStart, index)
        root.Right = dfs(preorder, preStart + index - postStart + 2, preEnd, postorder, index + 1, postEnd - 1)
        return root
    }
    return dfs(preorder, 0, n - 1, postorder, 0, n - 1)
}
