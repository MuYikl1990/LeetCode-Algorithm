package main

import "fmt"

func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 0
	sum := 0
	preNums := make(map[int]int)
	preNums[0] = 1
	for _, num := range nums {
		sum += num
		if count, ok := preNums[sum-k]; ok {
			res += count
		}
		preNums[sum]++
	}
	return res
}

func main() {
	nums, k := []int{1,2,6,4,-1,5}, 3
	res := subarraySum(nums, k)
	fmt.Println(res)
}
