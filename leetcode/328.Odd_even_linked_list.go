package main

type ListNode328 struct {
    Val int
    Next *ListNode328
}

func oddEvenList(head *ListNode328) *ListNode328 {
    if head == nil {
        return head
    }
    odd, even := head, head.Next
    evenHead := even
    for odd.Next != nil && even.Next != nil {
        odd.Next = odd.Next.Next
        even.Next = even.Next.Next
        odd, even = odd.Next, even.Next
    }
    odd.Next = evenHead
    return head
}

func main() {
	
}
