package main

import "fmt"

type MyQueue struct {
	stack1 []int
	stack2 []int
}


func Constructor232() MyQueue {
	return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
	this.stack1 = append(this.stack1, x)
}


func (this *MyQueue) Pop() int {
	if len(this.stack2) == 0 {
		for len(this.stack1) != 0 {
			this.stack2 = append([]int{this.stack1[len(this.stack1) - 1]}, this.stack2...)
			this.stack1 = this.stack1[:len(this.stack1) - 1]
		}
	}
	head := this.stack2[0]
	this.stack2 = this.stack2[1:]
	return head
}


func (this *MyQueue) Peek() int {
	if len(this.stack2) == 0 {
		for len(this.stack1) != 0 {
			this.stack2 = append([]int{this.stack1[len(this.stack1) - 1]}, this.stack2...)
			this.stack1 = this.stack1[:len(this.stack1) - 1]
		}
	}
	return this.stack2[0]
}


func (this *MyQueue) Empty() bool {
	if len(this.stack1) == 0 && len(this.stack2) == 0 {
		return true
	}
	return false
}

func main() {
	obj := Constructor232()
	x, y := 1, 2
	obj.Push(x)
	obj.Push(y)
	param1 := obj.Peek()
	param2 := obj.Pop()
	param3 := obj.Empty()
	fmt.Println(param1)
	fmt.Println(param2)
	fmt.Println(param3)
}
