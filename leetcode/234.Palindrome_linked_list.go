package main

import "fmt"

type ListNode234 struct {
    Val int
    Next *ListNode234
}

func isPalindrome(head *ListNode234) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	pre, cur := &ListNode234{}, &ListNode234{}
	for fast != nil && fast.Next != nil {
		cur = slow
		slow = slow.Next
		fast = fast.Next.Next
		cur.Next = pre
		pre = cur
	}

	if fast != nil {
		slow = slow.Next
	}

	for cur != nil && slow != nil {
		if cur.Val != slow.Val {
			return false
		}
		cur = cur.Next
		slow = slow.Next
	}
	return true
}

func main() {
	head := &ListNode234{Val: 1, Next: &ListNode234{Val: 2, Next: &ListNode234{Val: 2, Next: &ListNode234{Val: 1}}}}
	res := isPalindrome(head)
	fmt.Println(res)
}
