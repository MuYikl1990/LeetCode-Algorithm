package main

import (
	"container/heap"
	"context"
	"fmt"
	"math"
	"sort"
	"sync"
)

func removeKdigits1(s string) int {
	if len(s) == 0 {
		return 0
	}

	nFlag, sFlag, sign, num := false, false, 1, 0
	for _, char := range s {
		if char == ' ' && !nFlag && !sFlag {
			continue
		} else if (char == '+' || char == '-') && !nFlag && !sFlag {
			sFlag = true
			if char == '-' {
				sign = -1
			}
		} else if char >= '0' && char <= '9' {
			if num > math.MaxInt32 / 10 || (num == math.MaxInt32 / 10 && int(char - '0') > math.MaxInt32 % 10) {
				return math.MaxInt32
			} else if num < math.MinInt32 / 10 || (num == math.MinInt32 / 10 && int(char - '0') > 0) {
				return math.MinInt32
			}
			num = num * 10 + sign * int(char - '0')
			nFlag = true
		} else {
			break
		}
	}
	return num
}

func min11(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max11(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Do(ctx context.Context) {
	co := context.WithValue(ctx, "name", 789)
	c := co.Value("name1")
	fmt.Println(c)

}

type SliceQueue struct {
	data []interface{}

	sync.RWMutex
	m sync.Map
}

func (slice *SliceQueue) Locked(n int) {
	slice.RLock()
	defer slice.RUnlock()
	fgs := 789
	slice.m.Load(fgs)
}


type KthLargest struct {
	heap  []int
	k     int
}


func Constructor703(k int, nums []int) KthLargest {
	kth := KthLargest{
		heap: []int{},
		k :   k,
	}

	for i := range nums {
		kth.Add(nums[i])
	}
	return kth
}

func up(heap []int, idx int) {
	for {
		parent := (idx - 1) / 2
		if parent >= 0 && heap[idx] < heap[parent] {
			swap(heap, parent, idx)
		} else {
			break
		}
		idx = parent
	}
}

func down(heap []int, idx int) {
	n := len(heap)
	for {
		pos := idx
		if idx * 2 + 1 < n && heap[idx * 2 + 1] < heap[idx] {
			pos = idx * 2 + 1
		}
		if idx * 2 + 2 < n && heap[idx * 2 + 2] < heap[pos] {
			pos = idx * 2 + 2
		}
		if pos == idx {
			break
		}
		swap(heap, idx, pos)
		idx = pos
	}
}

func swap(nums []int, start, end int) {
	nums[start], nums[end] = nums[end], nums[start]
}


func (this *KthLargest) Add(val int) int {
	if len(this.heap) < this.k {
		this.heap = append(this.heap, val)
		up(this.heap, len(this.heap) - 1)
	} else {
		if this.heap[0] < val {
			this.heap[0] = val
			down(this.heap, 0)
		}
	}
	return this.heap[0]
}

func compressString(S string) string {
	n := len(S)
	left, right, res := 0, 0, ""
	for right < n {
		if right == n - 1 || S[left] != S[right] {
			length := right - left
			if S[left] == S[right] {
				length++
				right++
			}
			res += string(S[left])
			res += fmt.Sprintf("%d", length)
			left = right
		} else {
			right++
		}
	}
	if len(res) >= n {
		return S
	}
	return res
}

func reverseWords1(s string) string {
	s = " " + s
	n := len(s)
	res := ""
	left, right := n - 1, n
	for left >= 0 {
		if s[left] == ' ' {
			if left + 1 < right {
				res += s[left + 1:right] + " "
			}
			right = left
		}
		left--
	}
	return res[:len(res) - 1]
}

type MedianFinder1 struct {
	minHp Heap1
	maxHp MaxHeap
}

func Constructor1() MedianFinder1 {
	return MedianFinder1{}
}


func (this *MedianFinder1) AddNum(num int)  {
	minQ, maxQ := &this.minHp, &this.maxHp
	if minQ.Len() == maxQ.Len() {
		if maxQ.Len() == 0 || num <= minQ.IntSlice[0] {
			heap.Push(maxQ, num)
		} else {
			heap.Push(maxQ, heap.Pop(minQ))
			heap.Push(minQ, num)
		}
	} else {
		if num > maxQ.IntSlice[0] {
			heap.Push(minQ, num)
		} else {
			heap.Push(minQ, heap.Pop(maxQ))
			heap.Push(maxQ, num)
		}
	}
}


func (this *MedianFinder1) FindMedian() float64 {
	minQ, maxQ := &this.minHp, &this.maxHp
	fmt.Println(minQ, maxQ)
	if minQ.Len() != maxQ.Len() {
		return float64(maxQ.IntSlice[0])
	} else {
		return float64(maxQ.IntSlice[0] + minQ.IntSlice[0]) / 2
	}
}

type Heap1 struct {
	sort.IntSlice
}

func (hp *Heap1) Push(x interface{}) {
	hp.IntSlice = append(hp.IntSlice, x.(int))
}

func (hp *Heap1) Pop() interface{} {
	num := hp.IntSlice[len(hp.IntSlice) - 1]
	hp.IntSlice = hp.IntSlice[:len(hp.IntSlice) - 1]
	return num
}

type MaxHeap struct {
	sort.IntSlice
}

func (hp *MaxHeap) Less(i, j int) bool {
	return hp.IntSlice[i] > hp.IntSlice[j]
}

func (hp *MaxHeap) Push(x interface{}) {
	hp.IntSlice = append(hp.IntSlice, x.(int))
}

func (hp *MaxHeap) Pop() interface{} {
	num := hp.IntSlice[len(hp.IntSlice) - 1]
	hp.IntSlice = hp.IntSlice[:len(hp.IntSlice) - 1]
	return num
}

func main() {
	//num := "2147483646"
	//res := removeKdigits1(num)
	//fmt.Println(res)
	//ctx := context.TODO()
	//Do(ctx)
	//nums := []int{4,7,1,9,0,3,1,5,6,2}
	//ans := ba.HeapSort(nums)
	//fmt.Println(ans)
	//k, nums := 1, []int{}
	//obj := Constructor703(k, nums)
	//fmt.Println(obj.heap)
	//val := -3
	//param1 := obj.Add(val)
	//fmt.Println(param1)
	//val = -2
	//param1 = obj.Add(val)
	//fmt.Println(param1)
	//val = -4
	//param1 = obj.Add(val)
	//fmt.Println(param1)
	//val = 0
	//param1 = obj.Add(val)
	//fmt.Println(param1)
	//val = 4
	//param1 = obj.Add(val)
	//fmt.Println(param1)
	//fmt.Println(12)
	//fmt.Println(byte(12 / 10))
	//defer fmt.Print("1 in main\n")
	//defer func() {
	//	defer func() {
	//		fmt.Print("2\n")
	//		panic("panic again and again")
	//	}()
	//	fmt.Print("3\n")
	//	panic("panic again")
	//}()
	//s := "  a good   example"
	//res := reverseWords1(s)
	//fmt.Println(res)
	obj := Constructor1()
	obj.AddNum(1)
	obj.AddNum(2)
	param1 := obj.FindMedian()
	fmt.Println(param1)
	obj.AddNum(3)
	param2 := obj.FindMedian()
	fmt.Println(param2)
	obj.AddNum(1)
	param3 := obj.FindMedian()
	fmt.Println(param3)
}

