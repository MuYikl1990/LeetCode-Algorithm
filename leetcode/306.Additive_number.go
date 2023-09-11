package main

import "fmt"

func isAdditiveNumber(num string) bool {
	return dfs306(num, 0, 0, 0, 0)
}

func dfs306(num string, index, count, fir, sec int) bool {
	if index >= len(num) {
		return count > 2
	}

	cur := 0
	for i := index; i < len(num); i++ {
		if num[index] == '0' && i > index {
			return false
		}
		cur = cur * 10 + int(num[i] - '0')

		if count >= 2 {
			sum := fir + sec
			if cur > sum {
				return false
			}
			if cur < sum {
				continue
			}
		}

		if dfs306(num, i + 1, count + 1, sec, cur) {
			return true
		}
	}

	return false
}

func main() {
	num := "199100199"
	res := isAdditiveNumber(num)
	fmt.Println(res)
}
