package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	pre, cur := dummy, dummy.Next
	for i := 0; i < left - 1; i++ {
		pre, cur = pre.Next, cur.Next
	}

	for i := right - left; i > 0; i-- {
		remove := cur.Next
		cur.Next = cur.Next.Next
		remove.Next = pre.Next
		pre.Next = remove
	}
	return dummy.Next
}
