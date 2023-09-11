package main

import "fmt"

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}

	mid := start + (end - start) / 2
	res := mergeSort(nums, start, mid) + mergeSort(nums, mid + 1, end)
	i, j := start, mid + 1
	var tmp []int
	for i <= mid && j <= end {
		if nums[i] > nums[j] {
			tmp = append(tmp, nums[j])
			res += mid - i + 1
			j++
		} else {
			tmp = append(tmp, nums[i])
			i++
		}
	}

	tmp = append(tmp, nums[i:mid+1]...)
	tmp = append(tmp, nums[j:end+1]...)
	for k := start; k <= end; k++ {
		nums[k] = tmp[k - start]
	}
	return res
}

func main() {
	nums := []int{7,5,6,4}
	res := reversePairs(nums)
	fmt.Println(res)
}
