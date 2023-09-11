package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type MedianFinder struct {
	heapMin, heapMax heapQ
}

func Constructor295() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int)  {
	minQ, maxQ := &this.heapMin, &this.heapMax
	if maxQ.Len() == 0 || num <= -maxQ.IntSlice[0] {
		heap.Push(maxQ, -num)
		if minQ.Len() + 1 < maxQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	} else {
		heap.Push(minQ, num)
		if minQ.Len() > maxQ.Len() {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	minQ, maxQ := this.heapMin, this.heapMax
	if minQ.Len() == maxQ.Len() {
		return float64(minQ.IntSlice[0] - maxQ.IntSlice[0]) / 2
	}
	return float64(-maxQ.IntSlice[0])
}

type heapQ struct {
	sort.IntSlice
}

func (hp *heapQ) Push(x interface{}) {
	hp.IntSlice = append(hp.IntSlice, x.(int))
}

func (hp *heapQ) Pop() interface{} {
	number := hp.IntSlice[len(hp.IntSlice) - 1]
	hp.IntSlice = hp.IntSlice[:len(hp.IntSlice)-1]
	return number
}

func main() {
	obj := Constructor295()
	obj.AddNum(1)
	obj.AddNum(2)
	param1 := obj.FindMedian()
	fmt.Println(param1)
	obj.AddNum(3)
	param2 := obj.FindMedian()
	fmt.Println(param2)
}
