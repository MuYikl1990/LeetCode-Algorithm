package main

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp, length := head, 0
	for tmp != nil {
		length++
		tmp = tmp.Next
	}

	dummy := &ListNode{Val: 0}
	dummy.Next = head
	for i := 1; i < length; i <<= 1 {
		pre := dummy
		cur := dummy.Next
		for cur != nil {
			left := cur
			right := split(cur, i)
			cur = split(right, i)
			pre.Next = mergeNode(left, right)

			for pre.Next != nil {
				pre = pre.Next
			}
		}
	}
	return dummy.Next
}

func split(node *ListNode, step int) *ListNode {
	if node == nil {
		return node
	}

	for i := 1; i < step && node.Next != nil; i++ {
		node = node.Next
	}

	right := node.Next
	node.Next = nil
	return right
}

func mergeNode(left, right *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	cur := dummy
	for left != nil && right != nil {
		if left.Val <= right.Val {
			dummy.Next = left
			left = left.Next
		} else {
			dummy.Next = right
			right = right.Next
		}
		dummy = dummy.Next
	}
	if left != nil {
		dummy.Next = left
	} else {
		dummy.Next = right
	}
	return cur.Next
}
