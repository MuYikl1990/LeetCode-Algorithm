package main

type TreeNode102 struct {
    Val int
    Left *TreeNode102
    Right *TreeNode102
}

func levelOrder(root *TreeNode102) [][]int {
    var res [][]int
    var stack []*TreeNode102
    if root == nil {
        return res
    }
    stack = append(stack, root)

    for len(stack) > 0 {
        length := len(stack)
        var tmp []int
        for i := 0; i < length; i++ {
            cur := stack[i]
            tmp = append(tmp, cur.Val)
            if cur.Left != nil {
                stack = append(stack, cur.Left)
            }
            if cur.Right != nil {
                stack = append(stack, cur.Right)
            }
        }
        res = append(res, tmp)
        stack = stack[length:]
    }
    return res
}
