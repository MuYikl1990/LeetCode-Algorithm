package main

import (
	"fmt"
	"strconv"
)

func largestMultipleOfThree(digits []int) string {
	sum := 0
	var cnt [10]int
	var mod [3]int
	for _, digit := range digits {
		sum += digit
		cnt[digit]++
		mod[digit % 3]++
	}
	removeMod, removeCnt := 0, 0
	if sum % 3 == 1 {
		if mod[1] > 0 {
			removeMod, removeCnt = 1, 1
		} else {
			removeMod, removeCnt = 2, 2
		}
	} else if sum % 3 == 2 {
		if mod[2] > 0 {
			removeMod, removeCnt = 2, 1
		} else {
			removeMod, removeCnt = 1, 2
		}
	}

	tmp := ""
	for i := 0; i < 10; i++ {
		for j := 0; j < cnt[i]; j++ {
			if removeCnt > 0 && removeMod == i % 3 {
				removeCnt--
			} else {
				tmp += strconv.Itoa(i)
			}
		}
	}
	if tmp != "" && tmp[len(tmp) - 1] == '0' {
		return "0"
	}
	var res []byte
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[i])
	}
	return string(res)
}

func main() {
	digits := []int{8,6,7,1,0}
	res := largestMultipleOfThree(digits)
	fmt.Println(res)
}
