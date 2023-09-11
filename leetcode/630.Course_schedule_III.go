package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	hp := &Heap{}
	total := 0
	for _, course := range courses {
		if total + course[0] <= course[1] {
			total += course[0]
			heap.Push(hp, course[0])
		} else if hp.Len() > 0 && course[0] < hp.IntSlice[0] {
			total += course[0] - hp.IntSlice[0]
			heap.Pop(hp)
			heap.Push(hp, course[0])
		}
	}
	return hp.Len()
}

type Heap struct {
	sort.IntSlice
}

func (hp *Heap) Less(i, j int) bool {
	return hp.IntSlice[i] > hp.IntSlice[j]
}

func (hp *Heap) Push(x interface{}) {
	hp.IntSlice = append(hp.IntSlice, x.(int))
}

func (hp *Heap) Pop() interface{} {
	num := hp.IntSlice[len(hp.IntSlice) - 1]
	hp.IntSlice = hp.IntSlice[:len(hp.IntSlice) - 1]
	return num
}

func main() {
	courses := [][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}
	res := scheduleCourse(courses)
	fmt.Println(res)
}
