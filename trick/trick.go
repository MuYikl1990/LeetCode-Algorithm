package trick

import (
	"container/heap"
	"sort"
)

// 翻转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}


// 数字转字母
func intToChar(i int) rune {
	return rune('a' + i)
}


// 向上取整
func getTop(n, k int) int {
	return (n + k - 1) / k
}


// 正数原码等于自身
// 负数原码求补码: 除符号位外取反加1
// 原码：-5 -> 10000101, 补码：11111011
// 0 的原码与补码相同


// 不用加减乘除做加法
func add(a int, b int) int {
	for b != 0 {
		a, b = a ^ b, a & b << 1 // a ^ b 表示无进位和，a & b << 1 表示进位和
	}
	return a
}


// 快速幂算法
func pow(x, n int) int {
	res := 1
	for n > 0 {
		if n & 1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}


// 倍增乘法
func mul(a, k int) int {
	ans := 0
	for k > 0 {
		if k & 1 == 1 {
			ans += a
		}
		k >>= 1
		a += a
	}
	return ans
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 前序遍历（迭代）
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			cur := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			root = cur.Right
		}
	}
	return res
}


// 中序遍历（迭代）
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			cur := stack[len(stack) - 1]
			res = append(res, cur.Val)
			stack = stack[:len(stack) - 1]
			root = cur.Right
		}
	}
	return res
}


// 后序遍历（迭代）
func postorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	var pre *TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			cur := stack[len(stack) - 1]
			if cur.Right != nil && cur.Right != pre {
				root = cur.Right
			} else {
				res = append(res, cur.Val)
				stack = stack[:len(stack) - 1]
				pre = cur
				root = nil
			}
		}
	}
	return res
}


// Morris 遍历(模拟中序遍历)
func morris(root *TreeNode) []int {
	var res []int
	for root != nil {
		if root.Left != nil {
			cur := root.Left
			for cur.Right != nil && cur.Right != root {
				cur = cur.Right
			}

			if cur.Right == nil {
				cur.Right = root
				root = root.Left
			} else {
				res = append(res, cur.Val)
				cur.Right = nil
				root = root.Right
			}
		} else {
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}


// KMP 算法
func KMP(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if m < n {
		return -1
	}

	next := buildNext(needle)
	i, j := 0, 0
	for i < m && j < n {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == n {
		return i - j
	}
	return -1
}

func buildNext(needle string) []int {
	n := len(needle)
	next := make([]int, n + 1)
	next[0] = -1
	i, j := 1, 0
	for i < n {
		if j == -1 || needle[i] == needle[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}


// 约瑟夫环问题
func lastRemaining(n int, k int) int {
	idx := 0
	for i := 2; i <= n; i++ {
		idx = (idx + k) % i
	}
	return idx + 1
}


// TopK 问题
type KthLargest struct {
	heap  []int
	k     int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		heap: []int{},
		k :   k,
	}

	for i := range nums {
		kth.Add(nums[i])
	}
	return kth
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


// 堆用包实现，默认是小顶堆
type Heap struct {
	sort.IntSlice
}

// 大顶堆需要实现 Less 函数
func (hp *Heap) Less(i, j int) bool {
	return hp.IntSlice[i] > hp.IntSlice[j]
}

// 堆顶就是第一个节点(hp.IntSlice[0])
func (hp *Heap) Push(x interface{}) {
	hp.IntSlice = append(hp.IntSlice, x.(int))
}

// Pop 操作会先交换第一个节点和最后的节点，所以最后的节点才是最小/最大的
func (hp *Heap) Pop() interface{} {
	num := hp.IntSlice[len(hp.IntSlice) - 1]
	hp.IntSlice = hp.IntSlice[:len(hp.IntSlice) - 1]
	return num
}

func pushAndPopOfHeap(num int) {
	hp := &Heap{}
	_ = hp.IntSlice[0] // 最大值/最小值
	heap.Push(hp, num)
	heap.Pop(hp)
	hp.Len()
}
