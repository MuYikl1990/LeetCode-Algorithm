package main

import "fmt"

type ListNode61 struct {
    Val int
    Next *ListNode61
}

func rotateRight(head *ListNode61, k int) *ListNode61 {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	n, root := 1, head
	for root.Next != nil {
		root = root.Next
		n++
	}

	step := n - k % n
	if n == step {
		return head
	}

	root.Next = head
	for step > 0 {
		root = root.Next
		step--
	}
	res := root.Next
	root.Next = nil
	return res
}

func main() {
	head := &ListNode61{Val: 1, Next: &ListNode61{Val: 2, Next: &ListNode61{Val: 3, Next: &ListNode61{Val: 4, Next: &ListNode61{Val: 5}}}}}
	res := rotateRight(head, 2)
	fmt.Println(res)
}
