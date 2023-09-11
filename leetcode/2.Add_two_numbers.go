package main

type ListNode2 struct {
    Val int
    Next *ListNode2
}

func addTwoNumbers(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	res := &ListNode2{Val: -1}
	cur := res
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode2{Val: carry % 10}
		carry /= 10
		cur = cur.Next
	}

	return res.Next
}
