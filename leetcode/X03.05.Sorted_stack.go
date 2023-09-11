package main

import "fmt"

type SortedStack struct {
	stack1 []int
	stack2 []int
}

func Constructor03() SortedStack {
	return SortedStack{[]int{}, []int{}}
}

func (this *SortedStack) Push(val int)  {
	for len(this.stack1) > 0 {
		if this.stack1[len(this.stack1) - 1] >= val {
			break
		} else {
			this.stack2 = append(this.stack2, this.stack1[len(this.stack1) - 1])
			this.stack1 = this.stack1[:len(this.stack1) - 1]
		}
	}
	this.stack1 = append(this.stack1, val)
	for len(this.stack2) > 0 {
		this.stack1 = append(this.stack1, this.stack2[len(this.stack2) - 1])
		this.stack2 = this.stack2[:len(this.stack2) - 1]
	}
}

func (this *SortedStack) Pop()  {
	if len(this.stack1) > 0 {
		this.stack1 = this.stack1[:len(this.stack1) - 1]
	}
}

func (this *SortedStack) Peek() int {
	if len(this.stack1) > 0 {
		return this.stack1[len(this.stack1) - 1]
	}
	return -1
}

func (this *SortedStack) IsEmpty() bool {
	return len(this.stack1) == 0
}

func main() {
	obj := Constructor03()
	val := 1
	obj.Push(val)
	obj.Pop()
	param1 := obj.Peek()
	param2 := obj.IsEmpty()
	fmt.Println(param1)
	fmt.Println(param2)
}
