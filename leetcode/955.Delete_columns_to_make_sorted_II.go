package main

import "fmt"

func minDeletionSize(strs []string) int {
	res, delFlag := 0, false
	n, sLength := len(strs), len(strs[0])
	char := make([]bool, n)

	for i := 0; i < sLength; i++ {
		for j := 0; j < n - 1; j++ {
			delFlag = false
			if !char[j] && strs[j][i] > strs[j + 1][i] {
				res++
				delFlag = true
				break
			}
		}
		if delFlag {
			continue
		}
		for k := 0; k < n - 1; k++ {
			if strs[k][i] < strs[k + 1][i] {
				char[k] = true
			}
		}
	}
	return res
}

func main() {
	strs := []string{"xga","yfb","yfa"}
	res := minDeletionSize(strs)
	fmt.Println(res)
}
