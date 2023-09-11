package main

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	p := head
	for p != nil {
		cur := &Node{Val: p.Val}
		cur.Next = p.Next
		p.Next = cur
		p = p.Next.Next
	}

	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	dummy := &Node{Val: -1}
	p = head
	tmp := dummy
	for p != nil {
		tmp.Next = p.Next
		tmp = tmp.Next
		p.Next = tmp.Next
		p = p.Next
	}
	return dummy.Next
}
