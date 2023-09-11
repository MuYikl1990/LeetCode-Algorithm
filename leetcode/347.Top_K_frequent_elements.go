package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

// 计数排序
func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	numList := make([][]int, len(nums) + 1)
	for key := range numMap {
		numList[numMap[key]] = append(numList[numMap[key]], key)
	}

	var res []int
	for i := len(numList) - 1; i >= 0; i-- {
		res = append(res, numList[i]...)
		if len(res) >= k {
			break
		}
	}
	return res
}

// 快速排序
func topKFrequentWithQuick(nums []int, k int) []int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	var list [][]int
	for key, value := range numMap {
		list = append(list, []int{key, value})
	}

	res := make([]int, k)
	quickSort(list, 0, len(list) - 1, k, 0, res)
	return res
}

func quickSort(list [][]int, start, end, k, resIndex int, res []int) {
	rand.Seed(time.Now().UnixNano())
	swap := rand.Int() % (end - start + 1) + start
	list[start], list[swap] = list[swap], list[start]

	pivot, index := list[start][1], start
	for i := start + 1; i <= end; i++ {
		if list[i][1] >= pivot {
			list[index + 1], list[i] = list[i], list[index + 1]
			index++
		}
	}
	list[start], list[index] = list[index], list[start]

	if k <= index - start {
		quickSort(list, start, index - 1, k, resIndex, res)
	} else {
		for i := start; i <= index; i++ {
			res[resIndex] = list[i][0]
			resIndex++
		}

		if k > index - start + 1 {
			quickSort(list, index + 1, end, k - index + start - 1, resIndex, res)
		}
	}
}

// 堆排序
func topKFrequentWithStack(nums []int, k int) []int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	var list minH
	for key, value := range numMap {
		heap.Push(&list, [2]int{key, value})
		if list.Len() > k {
			heap.Pop(&list)
		}
	}

	var res []int
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&list).([2]int)[0])
	}
	return res
}

type minH [][2]int

func (h minH) Len() int {
	return len(h)
}

func (h minH) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h minH) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minH) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *minH) Pop() interface{} {
	num := (*h)[len(*h) - 1]
	*h = (*h)[:len(*h) - 1]
	return num
}

func main() {
	nums, k := []int{5,3,1,1,1,3,73,1}, 2
	res := topKFrequent(nums, k)
	fmt.Println(res)
}
