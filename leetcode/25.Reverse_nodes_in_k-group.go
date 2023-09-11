package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	count := 0
	for cur != nil && count != k {
		cur = cur.Next
		count++
	}

	if count == k {
		cur = reverseKGroup(cur, k)
		for count > 0 {
			tmp := head.Next
			head.Next = cur
			cur = head
			head = tmp
			count--
		}
		head = cur
	}
	return head
}
