package main

import (
	"fmt"
	"math"
)

func judgePoint24(cards []int) bool {
	nums := make([]float64, len(cards))
	for i := range cards {
		nums[i] = float64(cards[i])
	}
	return dfs679(nums)
}

func dfs679(nums []float64) bool {
	if len(nums) == 1 {
		return math.Abs(nums[0] - 24) < 1e-6
	}

	flag := false
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			n1, n2 := nums[i], nums[j]
			var newArr []float64
			for k := 0; k < len(nums); k++ {
				if k != i && k != j {
					newArr = append(newArr, nums[k])
				}
			}
			flag = flag || dfs679(append(newArr, n1 + n2))
			flag = flag || dfs679(append(newArr, n1 * n2))
			flag = flag || dfs679(append(newArr, n1 - n2))
			flag = flag || dfs679(append(newArr, n2 - n1))
			if n1 != 0 {
				flag = flag || dfs679(append(newArr, n2 / n1))
			}
			if n2 != 0 {
				flag = flag || dfs679(append(newArr, n1 / n2))
			}
			if flag {
				return true
			}
		}
	}
	return false
}

func main() {
	cards := []int{3,3,8,8}
	res := judgePoint24(cards)
	fmt.Println(res)
}
