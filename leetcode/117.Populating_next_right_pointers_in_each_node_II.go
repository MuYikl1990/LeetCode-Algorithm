package main

type Node117 struct {
    Val int
    Left *Node117
    Right *Node117
    Next *Node117
}

func connectII(root *Node117) *Node117 {
    cur := root
    for cur != nil {
        dummy := &Node117{Val: 101}
        pre := dummy
        for cur != nil {
            if cur.Left != nil {
                pre.Next = cur.Left
                pre = pre.Next
            }
            if cur.Right != nil {
                pre.Next = cur.Right
                pre = pre.Next
            }
            cur = cur.Next
        }
        cur = dummy.Next
    }
    return root
}
