package main

type Node116 struct {
    Val int
    Left *Node116
    Right *Node116
    Next *Node116
}

func connect(root *Node116) *Node116 {
    if root == nil {
        return nil
    }
    left := root.Left
    right := root.Right

    for left != nil {
        left.Next = right
        left = left.Right
        right = right.Left
    }
    connect(root.Left)
    connect(root.Right)
    return root
}

func connect1(root *Node116) *Node116 {
    if root == nil {
        return root
    }

    for left := root; left.Left != nil; left = left.Left {
        for node := left; node != nil; node = node.Next {
            node.Left.Next = node.Right

            if node.Next != nil {
                node.Right.Next = node.Next.Left
            }
        }
    }
    return root
}
