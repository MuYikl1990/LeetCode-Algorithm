package main

import "fmt"

// 摩尔投票
func majorityElementII(nums []int) []int {
	num1, count1 := nums[0], 0
	num2, count2 := nums[0], 0
	for _, num := range nums {
		if num == num1 {
			count1++
			continue
		}
		if num == num2 {
			count2++
			continue
		}
		if count1 == 0 {
			num1 = num
			count1++
			continue
		}
		if count2 == 0 {
			num2 = num
			count2++
			continue
		}
		count1--
		count2--
	}

	count1, count2 = 0, 0
	var res []int
	for _, num := range nums {
		if num == num1 {
			count1++
		} else if num == num2 {
			count2++
		}
	}
	if count1 > len(nums) / 3 {
		res = append(res, num1)
	}
	if count2 > len(nums) / 3 {
		res = append(res, num2)
	}
	return res
}

func main() {
	nums := []int{3,2,3}
	res := majorityElementII(nums)
	fmt.Println(res)
}
