package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	mapNum := make(map[int]struct{})
	preSum := make([]int, len(nums) + 1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i - 1] + nums[i - 1]
	}

	for i := 2; i < len(preSum); i++ {
		mapNum[(preSum[i - 2] % k)] = struct{}{}
		if _, ok := mapNum[(preSum[i] % k)]; ok {
			return true
		}
	}
	return false
}

func main() {
	nums, k := []int{23,2,6,4,7}, 13
	res := checkSubarraySum(nums, k)
	fmt.Println(res)
}
