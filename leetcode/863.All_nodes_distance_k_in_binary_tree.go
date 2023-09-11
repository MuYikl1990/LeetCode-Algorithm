package main

type TreeNode863 struct {
    Val int
    Left *TreeNode863
    Right *TreeNode863
}

func distanceK(root *TreeNode863, target *TreeNode863, k int) []int {
    parent := make(map[int]*TreeNode863)
    var res []int
    graph(root, parent)
    find(target, parent,nil, k, &res)
    return res
}

func graph(root *TreeNode863, parent map[int]*TreeNode863) {
    if root == nil {
        return
    }
    if root.Left != nil {
        parent[root.Left.Val] = root
        graph(root.Left, parent)
    }
    if root.Right != nil {
        parent[root.Right.Val] = root
        graph(root.Right, parent)
    }
}

func find(target *TreeNode863, parent map[int]*TreeNode863, from *TreeNode863, k int, res *[]int) {
    if target == nil {
        return
    }
    if k == 0 {
        *res = append(*res, target.Val)
        return
    }
    if target.Left != from {
        find(target.Left, parent, target, k - 1, res)
    }
    if target.Right != from {
        find(target.Right, parent, target, k - 1, res)
    }
    if parent[target.Val] != from {
        find(parent[target.Val], parent, target, k - 1, res)
    }
}
