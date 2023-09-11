package main

import "fmt"

func circularArrayLoop(nums []int) bool {
	var nextPos func(int) int
	nextPos = func(index int) int {
		tmp := index + nums[index]
		for tmp + len(nums) < 0 {
			tmp += len(nums)
		}
		return (tmp + len(nums)) % len(nums)
	}

	for i := range nums {
		slow := i
		fast := nextPos(slow)

		for nums[fast] * nums[i] > 0 && nums[nextPos(fast)] * nums[i] > 0 {
			if fast == slow {
				if slow == nextPos(slow) {
					break
				}
				return true
			}
			slow = nextPos(slow)
			fast = nextPos(nextPos(fast))
		}
	}
	return false
}

func main() {
	nums := []int{-2,1,-1,-2,-2}
	res := circularArrayLoop(nums)
	fmt.Println(res)
}
