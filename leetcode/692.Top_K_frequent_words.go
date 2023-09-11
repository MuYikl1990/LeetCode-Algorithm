package main

import (
	"container/heap"
	"fmt"
)

func topKFrequentWords(words []string, k int) []string {
	var stack minStack
	wordsMap := make(map[string]int)
	for i := range words {
		wordsMap[words[i]]++
	}

	for key, val := range wordsMap {
		heap.Push(&stack, element{key, val})
		if len(stack) > k {
			heap.Pop(&stack)
		}
	}

	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(&stack).(element).w
	}
	return res
}

type element struct {
	w string
	c int
}

type minStack []element

func (m minStack) Len() int {
	return len(m)
}

func (m minStack) Less(i, j int) bool {
	return m[i].c < m[j].c || (m[i].c == m[j].c && m[i].w > m[j].w)
}

func (m minStack) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minStack) Push(x interface{}) {
	*m = append(*m, x.(element))
}

func (m *minStack) Pop() interface{} {
	n := (*m)[len(*m) - 1]
	*m = (*m)[:len(*m) - 1]
	return n
}

func main() {
	words, k := []string{"i","love","leetcode","i","love","coding"}, 2
	res := topKFrequentWords(words, k)
	fmt.Println(res)
}
