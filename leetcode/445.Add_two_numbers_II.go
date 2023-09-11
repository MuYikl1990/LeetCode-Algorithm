package main

type ListNode445 struct {
    Val int
    Next *ListNode445
}

func addTwoNumbersII(l1 *ListNode445, l2 *ListNode445) *ListNode445 {
	var s1, s2 []int
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	carry := 0
	var res *ListNode445
	for len(s1) > 0 || len(s2) > 0 || carry != 0 {
		x, y := 0, 0
		if len(s1) > 0 {
			x = s1[len(s1) - 1]
			s1 = s1[:len(s1) - 1]
		}
		if len(s2) > 0 {
			y = s2[len(s2) - 1]
			s2 = s2[:len(s2) - 1]
		}
		cur := x + y + carry
		carry = cur / 10
		cur = cur % 10
		node := &ListNode445{Val: cur}
		node.Next = res
		res = node
	}
	return res
}
