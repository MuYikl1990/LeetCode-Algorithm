package main

type ListNode21 struct {
	Val int
	Next *ListNode21
}

func mergeTwoLists(list1 *ListNode21, list2 *ListNode21) *ListNode21 {
	head := &ListNode21{Val: -101}
	cur := head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			break
		}
		if list2 == nil {
			cur.Next = list1
			break
		}
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	return head.Next
}
