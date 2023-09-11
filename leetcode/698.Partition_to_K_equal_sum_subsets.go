package main

import (
	"fmt"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum % k != 0 {
		return false
	}
	target := sum / k
	sort.Ints(nums)
	if nums[len(nums) - 1] > target {
		return false
	}
	used := make([]bool, len(nums))
	return backtrace698(nums, len(nums) - 1, 0, k, target, used)
}

func backtrace698(nums []int, begin, sum, k, target int, used []bool) bool {
	if k == 1 {
		return true
	}
	if sum == target {
		return backtrace698(nums, len(nums) - 1, 0, k - 1, target, used)
	}
	for i := begin; i >= 0; i-- {
		if used[i] || sum + nums[i] > target {
			continue
		}
		used[i] = true
		if backtrace698(nums, i - 1, sum + nums[i], k, target, used) {
			return true
		}
		used[i] = false
		for i > 0 && nums[i - 1] == nums[i] {
			i--
		}
	}
	return false
}

func main() {
	nums, k := []int{4,3,2,3,5,2,5}, 4
	res := canPartitionKSubsets(nums, k)
	fmt.Println(res)
}
