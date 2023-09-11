package main

import "fmt"

func leastInterval(tasks []byte, n int) int {
	maxCnt, maxTask := 0, 0
	char := [26]int{}
	for _, task := range tasks {
		char[task - 'A']++
		maxCnt = max621(maxCnt, char[task - 'A'])
	}

	for i := 0; i < 26; i++ {
		if char[i] == maxCnt {
			maxTask++
		}
	}
	return max621(len(tasks), (maxCnt - 1) * (n + 1) + maxTask)
}

func max621(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tasks, n := []byte{'A','A','A','A','A','A','B','C','D','E','F','G'}, 2
	res := leastInterval(tasks, n)
	fmt.Println(res)
}
