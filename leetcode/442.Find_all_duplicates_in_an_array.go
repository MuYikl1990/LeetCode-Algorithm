package main

import "fmt"

func findDuplicates(nums []int) []int {
	var res []int
	if len(nums) < 2 {
		return res
	}

	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if nums[num - 1] > 0 {
			nums[num - 1] = -nums[num - 1]
		} else {
			res = append(res, num)
		}
	}
	return res
}

func main() {
	nums := []int{4,3,2,7,8,7,3,1}
	res := findDuplicates(nums)
	fmt.Println(res)
}
