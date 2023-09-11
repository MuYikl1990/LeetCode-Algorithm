package main

// 1->8->3->6->5->4->7->2->NULL
// 1. 按奇偶位置拆分链表，得1->3->5->7->NULL 和 8->6->4->2->NULL
// 2. 反转偶链表，得1->3->5->7->NULL 和 2->4->6->8->NULL
// 3. 合并两个有序链表，得1->2->3->4->5->6->7->8->NULL

type ListNode01 struct {
	Val int
	Next *ListNode01
}

func sortOddEvenList(head *ListNode01) *ListNode01 {
	if head == nil || head.Next == nil {
		return head
	}
	odd, even := head, head.Next
	evenHead := even
	for odd.Next != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd, even = odd.Next, even.Next
	}
	odd.Next = nil
	even = reverse02(evenHead)

	root := &ListNode01{Val: -1}
	node := root
	for odd != nil && even != nil {
		if odd.Val <= even.Val {
			node.Next = odd
			odd = odd.Next
		} else {
			node.Next = even
			even = even.Next
		}
		node = node.Next
	}
	if odd != nil {
		node.Next = odd
	}
	if even != nil {
		node.Next = even
	}
	return root.Next
}

// 递归
func reverse01(head *ListNode01) *ListNode01 {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverse01(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// 迭代
func reverse02(head *ListNode01) *ListNode01 {
	var pre *ListNode01
	cur, next := head, head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
