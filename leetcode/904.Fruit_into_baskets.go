package main

import "fmt"

func totalFruit(fruits []int) int {
	res := 0
	left, right, types := 0, 0, 0
	fMap := make(map[int]int)
	for right < len(fruits) {
		fMap[fruits[right]]++
		if fMap[fruits[right]] == 1 {
			types++
		}
		for types > 2 {
			fMap[fruits[left]]--
			if fMap[fruits[left]] == 0 {
				types--
			}
			left++
		}
		res = max904(res, right - left + 1)
		right++
	}
	return res
}

func max904(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fruits := []int{3,3,3,1,2,1,1,2,3,3,4}
	res := totalFruit(fruits)
	fmt.Println(res)
}
