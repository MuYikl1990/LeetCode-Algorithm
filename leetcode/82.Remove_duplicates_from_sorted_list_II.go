package main

type ListNode82 struct {
    Val int
    Next *ListNode82
}

func deleteDuplicates(head *ListNode82) *ListNode82 {
    if head == nil || head.Next == nil {
        return head
    }

    dummy := &ListNode82{Val: -101}
    dummy.Next = head
    pre := dummy
    cur := head
    for cur != nil {
        for cur.Next != nil && cur.Val == cur.Next.Val {
            cur = cur.Next
        }
        if pre.Next == cur {
            pre = pre.Next
        } else {
            pre.Next = cur.Next
        }
        cur = cur.Next
    }
    return dummy.Next
}
