package main

import (
	"container/heap"
	"fmt"
)

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	var stack maxStack
	res, fuel, distance, index := 0, startFuel, 0, 0
	for distance < target {
		if fuel == 0 {
			if len(stack) > 0 {
				fuel += heap.Pop(&stack).(int)
				res++
			} else {
				return -1
			}
		}
		distance += fuel
		fuel = 0
		for index < len(stations) && stations[index][0] <= distance {
			heap.Push(&stack, stations[index][1])
			index++
		}
	}
	return res
}

type maxStack []int

func (m maxStack) Len() int {
	return len(m)
}

func (m maxStack) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m maxStack) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxStack) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxStack) Pop() interface{} {
	n := (*m)[len(*m) - 1]
	*m = (*m)[:len(*m) - 1]
	return n
}

func main() {
	target, startFuel, stations := 100, 10, [][]int{{10,60}, {20, 30}, {30, 30}, {60, 40}}
	res := minRefuelStops(target, startFuel, stations)
	fmt.Println(res)
}
