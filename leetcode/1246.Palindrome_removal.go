package main

import (
	"fmt"
	"math"
)

// 给你一个整数数组 nums，每一次操作你都可以选择并删除它的一个回文子数组，每当你删除掉一个子数组，右侧元素都会自行向前移动填补空位
// 计算并返回从数组中删除所有数字所需的最少操作次数

// dp[i][j]表示删除arr[i]~arr[j]中所有数组所需要的最少操作数
// 1. i > j时, dp[i][j] == 0;
// 2. i == j时, 说明arr[i]就是arr[j], 删除自己本身最少需要操作一次, dp[i][j] == 1;
// 3. i + 1 == j时, 说明arr[i]和arr[j]之间相差一个字符, 那么如果arr[i],arr[j]构成回文则dp[i][j] == 1否则dp[i][j] == 2
// 4. i + 1 < j时, 说明arr[i]和arr[j]之间相差不止一个字符,则有两种情况:
//    a. arr[i+1]~arr[j-1]是回文串且arr[i] == arr[j], 说明arr[i]~arr[j]也是回文串
//       那么删除arr[i]~arr[j]和删除arr[i+1]~arr[j-1]实际上是一样的
//       此时, dp[i][j] = min(dp[i][j], dp[i+1][j-1])
//    b. arr[i+1]~arr[j-1]不是回文串, 那么此时不论arr[i]是否等于arr[j], arr[i]~arr[j]都一定可以拆分成两个部分,
//       即arr[i]~arr[k]和arr[k+1]~arr[j]使得删除arr[i]~arr[j]操作最少, dp[i][k] + dp[k+1][j]
//    a,b两种情况我们需要取min
func palindromeRemoval(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		for i := j; i >= 0; i-- {
			if i == j {
				dp[i][j] = 1
				continue
			}

			if i + 1 == j {
				dp[i][j] = 2
				if nums[i] == nums[j] {
					dp[i][j] = 1
				}
				continue
			}

			dp[i][j] = math.MaxInt
			if nums[i] == nums[j] {
				dp[i][j] = min1246(dp[i][j], dp[i + 1][j - 1])
			}

			for k := i; k < j; k++ {
				dp[i][j] = min1246(dp[i][j], dp[i][k] + dp[k + 1][j])
			}
		}
	}
	return dp[0][n - 1]
}

func min1246(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{1,3,4,3,1,5}
	res := palindromeRemoval(nums)
	fmt.Println(res)
}
