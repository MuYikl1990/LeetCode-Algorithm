package main

import "fmt"

// fast-low point
func exchange(nums []int) []int {
	odd, even := 0, 0
	for odd < len(nums) {
		if nums[odd] % 2 == 1 {
			if odd > even {
				nums[odd], nums[even] = nums[even], nums[odd]
			}
			even++
		}
		odd++
	}
	return nums
}

// head-tail point
func exchange1(nums []int) []int {
	even, odd := 0, len(nums) - 1
	for even < odd {
		if nums[odd] % 2 == 0 {
			odd--
			continue
		}
		if nums[even] % 2 == 1 {
			even++
			continue
		}
		nums[even], nums[odd] = nums[odd], nums[even]
		even++
		odd--
	}
	return nums
}

func main() {
	nums := []int{2,1,3,4}
	res := exchange1(nums)
	fmt.Println(res)
}
