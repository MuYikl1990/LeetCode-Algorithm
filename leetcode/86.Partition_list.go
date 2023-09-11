package main

import "fmt"

type ListNode86 struct {
    Val int
    Next *ListNode86
}

// in place
func partitionList(head *ListNode86, x int) *ListNode86 {
	if head == nil {
		return head
	}

	dummy := &ListNode86{Val: -101}
	dummy.Next = head
	left := dummy

	for left.Next != nil && left.Next.Val < x {
		left = left.Next
	}
	if left.Next == nil {
		return head
	}

	right := left
	for right.Next != nil {
		if right.Next.Val < x {
			cur := right.Next
			leftRight := left.Next
			rightRight := cur.Next
			left.Next = cur
			cur.Next = leftRight
			right.Next = rightRight
			left = left.Next
		} else {
			right = right.Next
		}
	}
	return dummy.Next
}

// two dummy join
func partitionList1(head *ListNode86, x int) *ListNode86 {
	if head == nil {
		return head
	}

	small, large := &ListNode86{Val: -101}, &ListNode86{Val: 101}
	s, l := small, large

	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}

	large.Next = nil
	small.Next = l.Next
	return s.Next
}

func main() {
	head, x := &ListNode86{Val: 1, Next: &ListNode86{Val: 4, Next: &ListNode86{Val: 3, Next: &ListNode86{Val: 2, Next: &ListNode86{Val: 5, Next: &ListNode86{Val: 2}}}}}}, 3
	res := partitionList(head, x)
	fmt.Println(res.Next.Next.Next.Next.Next)
}
