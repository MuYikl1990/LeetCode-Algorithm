package main

import "fmt"

func majorityElement(nums []int) int {
	res, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			res = nums[i]
		}
		if nums[i] == res {
			count++
		} else {
			count--
		}
	}
	return res
}

func main() {
	nums := []int{2,2,1,1,1,2,2}
	res := majorityElement(nums)
	fmt.Println(res)
}
