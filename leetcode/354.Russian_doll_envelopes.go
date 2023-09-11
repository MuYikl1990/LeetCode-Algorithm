package main

import (
	"fmt"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})

	// spider cards
	piles := 0
	stack := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		left, right := 0, piles
		height := envelopes[i][1]
		for left < right {
			mid := left + (right - left) / 2
			if stack[mid] >= height {
				right = mid
			} else {
				left = mid + 1
			}
		}

		if left == piles {
			piles++
		}
		stack[right] = height
	}
	return piles
}

func main() {
	envelopes := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	//envelopes = [][]int{{1,1},{1,1},{1,1}}
	res := maxEnvelopes(envelopes)
	fmt.Println(res)
}
