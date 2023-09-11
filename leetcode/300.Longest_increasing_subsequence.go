package main

import "fmt"

func lengthOfLIS(nums []int) int {
	stack := make([]int, len(nums))
	piles := 0

	for i := 0; i < len(nums); i++ {
		left, right, num := 0, piles, nums[i]
		for left < right {
			mid := left + (right - left) / 2
			if stack[mid] >= num {
				right = mid
			} else {
				left = mid + 1
			}
		}

		if left == piles {
			piles++
		}
		stack[left] = num
	}
	return piles
}

func main() {
	nums := []int{10,9,2,5,3,7,101,18}
	res := lengthOfLIS(nums)
	fmt.Println(res)
}
