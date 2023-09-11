package main

import (
	"container/heap"
	"fmt"
)

func furthestBuilding(heights []int, bricks int, ladders int) int {
	length, sum := len(heights), 0
	var stack minQ
	for i := 1; i < length; i++ {
		if heights[i - 1] < heights[i] {
			heap.Push(&stack, heights[i] - heights[i - 1])
			if stack.Len() > ladders {
				min := heap.Pop(&stack)
				sum += min.(int)
				if sum > bricks {
					return i - 1
				}
			}
		}
	}
	return length - 1
}

type minQ []int

func (m minQ) Len() int {
	return len(m)
}

func (m minQ) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m minQ) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minQ) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minQ) Pop() interface{} {
	n := (*m)[len(*m) - 1]
	*m = (*m)[:len(*m) - 1]
	return n
}

func main() {
	heights, bricks, ladders := []int{4,12,2,7,3,18,20,3,19}, 10, 2
	res := furthestBuilding(heights, bricks, ladders)
	fmt.Println(res)
}
