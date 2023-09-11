package main

import "fmt"

func minimumJumps(forbidden []int, a int, b int, x int) int {
	if x == 0 {
		return 0
	}

	res := -1
	banMap := make(map[int]struct{})
	for i := range forbidden {
		banMap[forbidden[i]] = struct{}{}
	}
	var dfs func(int, int, int)
	dfs = func(distance, count, flag int) {
		if 0 <= distance && distance <= 6000 {
			if distance == x {
				res = count
			}
			if _, ok := banMap[distance + a]; !ok {
				banMap[distance + a] = struct{}{}
				dfs(distance + a, count + 1, 0)
			}
			if _, ok := banMap[distance - b]; !ok && flag == 0 {
				dfs(distance - b, count + 1, 1)
			}
		}
	}
	dfs(0, 0, 0)
	return res
}

func main() {
	forbidden, a, b, x := []int{1,6,2,14,5,17,4}, 16, 9, 7
	res := minimumJumps(forbidden, a, b, x)
	fmt.Println(res)
}
