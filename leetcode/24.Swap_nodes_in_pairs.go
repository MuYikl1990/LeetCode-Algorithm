package main

type ListNode24 struct {
	Val int
	Next *ListNode24
}

func swapPairs(head *ListNode24) *ListNode24 {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}
