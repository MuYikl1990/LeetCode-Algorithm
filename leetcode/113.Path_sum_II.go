package main

type TreeNode113 struct {
    Val int
    Left *TreeNode113
    Right *TreeNode113
}

func pathSumII(root *TreeNode113, targetSum int) [][]int {
    var res [][]int
    var dfs func(*TreeNode113, []int, int)
    dfs = func(root *TreeNode113, tmp []int, targetSum int) {
        if root == nil {
            return
        }
        if root.Left == nil && root.Right == nil {
            if root.Val == targetSum {
                ans := make([]int, len(tmp))
                copy(ans, tmp)
                res = append(res, ans)
            }
            return
        }
        tmp = append(tmp, root.Val)
        dfs(root.Left, tmp, targetSum - root.Val)
        dfs(root.Right, tmp, targetSum - root.Val)
        tmp = tmp[:len(tmp) - 1]
        return
    }
    dfs(root, []int{}, targetSum)
    return res
}
