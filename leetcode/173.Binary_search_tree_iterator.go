package main

type TreeNode173 struct {
    Val int
    Left *TreeNode173
    Right *TreeNode173
}

type BSTIterator struct {
	stack []*TreeNode173
}


func Constructor173(root *TreeNode173) BSTIterator {
	var res BSTIterator
	for root != nil {
		res.stack = append(res.stack, root)
		root = root.Left
	}
	return res
}


func (this *BSTIterator) Next() int {
	length := len(this.stack)
	cur := this.stack[length - 1]
	this.stack = this.stack[:length - 1]
	node := cur.Right
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
	return cur.Val
}


func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
