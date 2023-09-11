package main

import "fmt"

type ListNode83 struct {
	Val int
	Next *ListNode83
}

func deleteDuplicatesI(head *ListNode83) *ListNode83 {
	if head == nil {
		return head
	}

	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func main() {
	head := &ListNode83{Val: 1, Next: &ListNode83{Val: 1, Next: &ListNode83{Val: 2, Next: &ListNode83{Val: 2, Next: &ListNode83{Val: 3}}}}}
	res := deleteDuplicatesI(head)
	fmt.Println(res)
}
