package main

import (
	"fmt"
	"math/rand"
)

type ListNode382 struct {
    Val int
    Next *ListNode382
}

type Solution382 struct {
	head *ListNode382
}

func Constructor382(head *ListNode382) Solution382 {
	return Solution382{head}
}

// 蓄水池抽样算法
func (this *Solution382) GetRandom() int {
	res, i := 0, 1
	node := this.head

	for node != nil {
		if rand.Intn(i) == 0 {
			res = node.Val
		}
		i++
		node = node.Next
	}
	return res
}

func main() {
	head := &ListNode382{Val: 2, Next: &ListNode382{Val: 5, Next: &ListNode382{Val: 8}}}
	obj := Constructor382(head)
	param1 := obj.GetRandom()
	fmt.Println(param1)
	param2 := obj.GetRandom()
	fmt.Println(param2)
	param3 := obj.GetRandom()
	fmt.Println(param3)
}
