// 给定二叉树其中的一个结点，请找出其中序遍历顺序的下一个结点并且返回。
// 注意，树中的结点不仅包含左右子结点，而且包含指向父结点的指针。

package main

type TreeLinkNode struct {
	Val int
	Left *TreeLinkNode
	Right *TreeLinkNode
	Next *TreeLinkNode
}

func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode.Right != nil {
		lNode := pNode.Right
		for lNode.Left != nil {
			lNode = lNode.Left
		}
		return lNode
	}
	p := pNode
	for p.Next != nil {
		if p.Next.Left == p {
			return p.Next
		}
		p = p.Next
	}
	return nil
}
