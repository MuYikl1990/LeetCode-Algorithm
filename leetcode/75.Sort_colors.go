package main

import "fmt"

func sortColors(nums []int)  {
	n0, n1 := 0, 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		nums[i] = 2
		if num < 2 {
			nums[n1] = 1
			n1++
		}
		if num < 1 {
			nums[n0] = 0
			n0++
		}
	}
}

func main() {
	nums := []int{2,0,2,1,1,0}
	sortColors(nums)
	fmt.Println(nums)
}
