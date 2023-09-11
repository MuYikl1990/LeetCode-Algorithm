package main

import (
	"fmt"
	"math/rand"
)

const (
	MaxLevel = 32
	P = 0.25
)

type Skiplist struct {
	head *SkipNode
	level int
}

type SkipNode struct {
	val int
	next []*SkipNode
}

func Constructor1206() Skiplist {
	return Skiplist{head: &SkipNode{-1, make([]*SkipNode, MaxLevel)}, level: 0}
}

func (this *Skiplist) Search(target int) bool {
	level := this.level
	cur := this.head
	for i := level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < target {
			cur = cur.next[i]
		}
	}
	cur = cur.next[0]
	return cur != nil && cur.val == target
}

func (this *Skiplist) Add(num int)  {
	update := make([]*SkipNode, MaxLevel)
	for i := range update {
		update[i] = this.head
	}
	cur := this.head
	level := this.level
	for i := level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		update[i] = cur
	}
	lv := this.randomLevel()
	this.level = max1206(this.level, lv)
	newNode := &SkipNode{val: num, next: make([]*SkipNode, lv)}
	for i, node := range update[:lv] {
		newNode.next[i] = node.next[i]
		node.next[i] = newNode
	}
}

func (this *Skiplist) Erase(num int) bool {
	level := this.level
	update := make([]*SkipNode, level)
	for i := range update {
		update[i] = this.head
	}
	cur := this.head
	for i := level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		update[i] = cur
	}
	cur = cur.next[0]
	if cur == nil || cur.val != num {
		return false
	}

	for i := 0; i < level && update[i].next[i] == cur; i++ {
		update[i].next[i] = cur.next[i]
	}

	for this.level > 1 && this.head.next[this.level - 1] == nil {
		this.level--
	}
	return true
}

func (this *Skiplist) randomLevel() int {
	level := 1
	for level < MaxLevel && rand.Float32() < P {
		level++
	}
	return level
}

func max1206(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	target, num1, num2, num3, num4, num5 := 0, 1, 2, 3, 4, 5
	obj := Constructor1206()
	param := obj.Search(target)
	fmt.Println(param)
	obj.Add(num1)
	obj.Add(num2)
	obj.Add(num3)
	obj.Add(num4)
	obj.Add(num5)
	param = obj.Erase(num3)
	fmt.Println(param)
}
