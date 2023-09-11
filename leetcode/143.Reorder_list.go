package main

type ListNode143 struct {
    Val int
    Next *ListNode143
}

func reorderList(head *ListNode143)  {
    if head == nil || head.Next == nil {
        return
    }

    // find middle node
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // reverse half nodes
    node := slow.Next
    newHead := reverseTreeNode(node)
    slow.Next = nil

    // merge nodes
    tmp := head
    for tmp != nil && newHead != nil {
        pre := tmp.Next
        cur := newHead.Next
        tmp.Next = newHead
        newHead.Next = pre
        tmp = pre
        newHead = cur
    }
}

func reverseTreeNode(head *ListNode143) *ListNode143 {
    if head == nil || head.Next == nil {
        return head
    }

    cur := reverseTreeNode(head.Next)
    head.Next.Next = head
    head.Next = nil
    return cur
}
