package main

func hasCycle(head *ListNode) bool {
	//if head == nil || head.Next == nil {
	//	return false
	//}
	//
	//slow, fast := head, head.Next
	//for slow != fast {
	//	if fast == nil || fast.Next == nil {
	//		return false
	//	}
	//	slow = slow.Next
	//	fast = fast.Next.Next
	//}
	//return true
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
