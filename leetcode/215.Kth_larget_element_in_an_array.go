package main

import (
	"fmt"
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	length := len(nums)
	target := length - k
	left, right := 0, length - 1
	rand.Seed(time.Now().UnixNano())
	p := rand.Int() % length
	nums[left], nums[p] = nums[p], nums[left]

	for {
		index := partition(nums, left, right)
		if index == target {
			return nums[index]
		} else if index < target {
			left = index + 1
		} else {
			right = index - 1
		}
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left

	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func main() {
	nums := []int{3,2,1,5,6,4}
	res := findKthLargest(nums, 3)
	fmt.Println(res)
}
