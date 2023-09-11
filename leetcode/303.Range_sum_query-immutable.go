package main

import "fmt"

type NumArray struct {
	prefix []int
}

func Constructor303(nums []int) NumArray {
	prefix := make([]int, len(nums) + 1)
	for i := range nums {
		prefix[i + 1] = prefix[i] + nums[i]
	}
	return NumArray{prefix}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefix[right + 1] - this.prefix[left]
}

func main() {
	nums, left, right := []int{-2, 0, 3, -5, 2, -1}, 0, 2
	obj := Constructor303(nums)
	param1 := obj.SumRange(left,right)
	fmt.Println(param1)
	left, right = 2, 5
	param2 := obj.SumRange(left,right)
	fmt.Println(param2)
	left, right = 0, 5
	param3 := obj.SumRange(left,right)
	fmt.Println(param3)
}
