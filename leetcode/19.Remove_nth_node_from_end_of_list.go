package main

type ListNode19 struct {
    Val int
    Next *ListNode19
}

func removeNthFromEnd(head *ListNode19, n int) *ListNode19 {
	res := &ListNode19{}
	slow, fast := head, head
	res.Next = slow
	for n > 0 {
		fast = fast.Next
		n--
	}
	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return res.Next
}
