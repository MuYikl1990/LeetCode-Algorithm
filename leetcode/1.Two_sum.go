package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var res []int
	if len(nums) == 0 {
		return res
	}

	sumMap := make(map[int]int)
	for i, num := range nums {
		if val, ok := sumMap[target-num]; ok {
			res = []int{val, i}
			return res
		}
		sumMap[num] = i
	}
	return res
}

func main() {
	nums, target := []int{3,2,4}, 6
	res := twoSum(nums, target)
	fmt.Println(res)
}
