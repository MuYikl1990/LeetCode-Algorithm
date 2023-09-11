package main

import (
	"fmt"
	"math"
)

// 状压dp
func numSquarefulPerms(nums []int) int {
	n := len(nums)
	dp := make([][]int, 1 << n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(mask int, k int) int {
		if mask == 0 {
			return 1
		}

		if dp[mask][k] > -1 {
			return dp[mask][k]
		}

		count := 0
		for i := 0; i < n; i++ {
			if (mask >> i) & 1 == 1 && isSquare(nums[k] + nums[i]) {
				count += dfs(mask ^ (1 << i), i)
			}
		}
		dp[mask][k] = count
		return count
	}

	res, mask := 0, (1 << n) - 1
	for j := 0; j < n; j++ {
		res += dfs(mask ^ (1 << j), j)
	}

	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
	}

	for _, val := range cnt {
		res /= fac(val)
	}
	return res
}

func fac(n int) int {
	mul := 1
	for i := 2; i <= n; i++ {
		mul *= i
	}
	return mul
}

func isSquare(sum int) bool {
	x := int(math.Sqrt(float64(sum)))
	return x * x == sum
}

func main() {
	nums := []int{1, 17, 8}
	res := numSquarefulPerms(nums)
	fmt.Println(res)
}
