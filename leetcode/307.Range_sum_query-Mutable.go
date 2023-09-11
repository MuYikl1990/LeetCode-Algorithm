package main

import "fmt"

type NumArray307 struct {
	nums   []int
	preSum []int
}


func Constructor307(nums []int) NumArray307 {
	preSum := make([]int, len(nums) + 1)
	for i := range nums {
		preSum[i + 1] = preSum[i] + nums[i]
	}
	return NumArray307{nums: nums, preSum: preSum}
}


func (this *NumArray307) Update(index int, val int)  {
	diff := val - this.nums[index]
	this.nums[index] = val
	for i := index + 1; i < len(this.preSum); i++ {
		this.preSum[i] += diff
	}
}


func (this *NumArray307) SumRange(left int, right int) int {
	return this.preSum[right + 1] - this.preSum[left]
}

func main() {
	nums, index, val := []int{1, 3, 5}, 1, 2
	left, right := 0 ,2
	obj := Constructor307(nums)
	obj.Update(index,val)
	param := obj.SumRange(left,right)
	fmt.Println(param)
}
