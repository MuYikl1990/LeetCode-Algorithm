package main

type ListNode22 struct {
    Val int
    Next *ListNode22
}

func getKthFromEnd(head *ListNode22, k int) *ListNode22 {
	left, right := head, head
	for k > 0 {
		right = right.Next
		k--
	}
	for right != nil {
		left = left.Next
		right = right.Next
	}
	return left
}
