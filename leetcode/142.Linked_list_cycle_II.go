package main

type ListNode142 struct {
	Val  int
	Next *ListNode142
}

func detectCycle(head *ListNode142) *ListNode142 {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}
	return nil
}


