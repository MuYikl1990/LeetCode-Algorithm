package main

func mergeKLists(lists []*ListNode) *ListNode {
	return merge23(lists, 0, len(lists) - 1)
}

func merge23(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	} else if left > right {
		return nil
	}
	mid := (left + right) / 2
	return mergeTwoNode(merge23(lists, left, mid), merge23(lists, mid + 1, right))
}

func mergeTwoNode(node1, node2 *ListNode) *ListNode {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	dummy := &ListNode{Val: 0}
	head := dummy
	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			head.Next = node1
			node1 = node1.Next
		} else {
			head.Next = node2
			node2 = node2.Next
		}
		head = head.Next
	}
	if node1 != nil {
		head.Next = node1
	}
	if node2 != nil {
		head.Next = node2
	}
	return dummy.Next
}
