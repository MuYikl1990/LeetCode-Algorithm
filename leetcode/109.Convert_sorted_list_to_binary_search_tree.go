package main

type ListNode109 struct {
    Val int
    Next *ListNode109
}

type TreeNode109 struct {
    Val int
    Left *TreeNode109
    Right *TreeNode109
}

func sortedListToBST(head *ListNode109) *TreeNode109 {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	var pre *ListNode109

	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	root := &TreeNode109{Val: slow.Val}
	if pre != nil {
		pre.Next = nil
		root.Left = sortedListToBST(head)
	}
	root.Right = sortedListToBST(slow.Next)
	return root
}
