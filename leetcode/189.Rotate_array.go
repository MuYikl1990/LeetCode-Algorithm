package main

import "fmt"

func rotate189(nums []int, k int)  {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums) - 1; i < j; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	nums := []int{1,2,3,4,5,6,7}
	k := 3
	rotate189(nums, k)
	fmt.Println(nums)
}
