package main

import "fmt"

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	cur := 0
	for cur != slow {
		slow = nums[slow]
		cur = nums[cur]
	}
	return cur
}

func main() {
	nums := []int{1,3,4,2,2}
	res := findDuplicate(nums)
	fmt.Println(res)
}
