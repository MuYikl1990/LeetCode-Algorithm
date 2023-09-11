package main

import (
	"fmt"
)

type MinStack struct {
	stack []int
	min   int
}

func Constructor155() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int)  {
	if len(this.stack) == 0 {
		this.min = val
	} else {
		if val <= this.min {
			this.stack = append(this.stack, this.min)
			this.min = val
		}
	}
	this.stack = append(this.stack, val)
}


func (this *MinStack) Pop()  {
	if len(this.stack) != 0 {
		remove := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		if remove == this.min && len(this.stack) != 0 {
			this.min = this.stack[len(this.stack)-1]
			this.stack = this.stack[:len(this.stack)-1]
		}
	}
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}


func (this *MinStack) GetMin() int {
	return this.min
}

func main() {
	obj := Constructor155()
	obj.Push(2)
	obj.Push(0)
	obj.Push(3)
	obj.Push(-1)
	obj.Pop()
	param1 := obj.Top()
	fmt.Println(param1)
	param2 := obj.GetMin()
	fmt.Println(param2)
}
