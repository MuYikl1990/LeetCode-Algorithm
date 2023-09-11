package main

type TreeNode26 struct {
    Val int
    Left *TreeNode26
    Right *TreeNode26
}

func isSubStructure(A *TreeNode26, B *TreeNode26) bool {
    if A == nil || B == nil {
        return false
    }
    var isSub func(*TreeNode26, *TreeNode26) bool
    isSub = func(A *TreeNode26, B *TreeNode26) bool {
        if B == nil {
            return true
        }
        if A == nil || A.Val != B.Val {
            return false
        }
        return isSub(A.Left, B.Left) && isSub(A.Right, B.Right)
    }
    return isSub(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
