package main

import (
	"fmt"
	"math"
)

// hashmap
func longestConsecutive(nums []int) int {
	contain, res := make(map[int]int), 0
	for _, num := range nums {
		contain[num] = num
	}

	for i := range nums {
		cur := nums[i]
		if _, ok := contain[cur - 1]; !ok {
			right := contain[cur]
			for true {
				if val, ok := contain[right + 1]; ok {
					right = val
				} else {
					break
				}
			}
			contain[cur] = right
			res = max128(res, right - cur + 1)
		}
	}
	return res
}

// dp
func longestConsecutive1(nums []int) int {
	contain, res := make(map[int]int), 0

	for i := range nums {
		cur := nums[i]
		if _, ok := contain[cur]; !ok {
			left, right := 0, 0
			if val, ok := contain[cur - 1]; ok {
				left = val
			}
			if val, ok := contain[cur + 1]; ok {
				right = val
			}
			length := left + right + 1
			res = max128(res, length)
			contain[cur] = -1
			contain[cur - left] = length
			contain[cur + right] = length
		}
	}
	return res
}

// union find set
func longestConsecutive2(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	parent, count := make(map[int]int), make(map[int]int)
	for i := range nums {
		parent[nums[i]] = nums[i]
		count[nums[i]] = 1
	}
	res := 1
	for i := range nums {
		if find128(nums[i] + 1, parent) >= int(math.Pow(-10, 9)) && find128(nums[i] + 1, parent) <= int(math.Pow(10, 9)) {
			res = max128(res, union128(nums[i], nums[i] + 1, parent, count))
		}
	}
	return res
}

func find128(x int, parent map[int]int) int {
	if _, ok := parent[x]; !ok {
		return int(math.Pow(10, 9)) + 1
	}

	for parent[x] != x {
		parent[x] = parent[parent[x]]
		x = parent[x]
	}
	return x
}

func union128(x, y int, parent, count map[int]int) int {
	rootX := find128(x, parent)
	rootY := find128(y, parent)

	if rootX == rootY {
		return count[rootX]
	}
	parent[rootX] = rootY
	count[rootY] = count[rootX] + count[rootY]
	return count[rootY]
}

func max128(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{0,3,7,2,5,8,4,6,0,1}
	res := longestConsecutive(nums)
	fmt.Println(res)
}
