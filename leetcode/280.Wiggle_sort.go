package main

import "fmt"

func wiggleSortI(nums []int) {
	if len(nums) < 2 {
		return
	}

	for i := 1; i < len(nums); i++ {
		if (i % 2 == 0 && nums[i - 1] < nums[i]) || (i % 2 == 1 && nums[i - 1] > nums[i]) {
			nums[i - 1], nums[i] = nums[i], nums[i - 1]
		}
	}
}

func main() {
	nums := []int{3,5,2,1,6,4}
	wiggleSortI(nums)
	fmt.Println(nums)
}
