package main

import "fmt"

func change(amount int, coins []int) int {
	dp := make([]int, amount + 1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i - coin]
		}
	}
	return dp[amount]
}

func main() {
	amount, coins := 5, []int{1, 2, 5}
	res := change(amount, coins)
	fmt.Println(res)
	fmt.Println(string('a' + 27))
}
