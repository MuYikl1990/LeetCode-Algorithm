package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	var res, queue []int
	for right := range nums {
		for len(queue) != 0 && nums[queue[len(queue)-1]] <= nums[right] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, right)
		left := right - k + 1
		if queue[0] < left {
			queue = queue[1:]
		}
		if right + 1 >= k {
			res = append(res, nums[queue[0]])
		}
	}
	return res
}

func main() {
	nums, k := []int{4,3,5,4,3,3,6,7}, 3
	res := maxSlidingWindow(nums, k)
	fmt.Println(res)
}
