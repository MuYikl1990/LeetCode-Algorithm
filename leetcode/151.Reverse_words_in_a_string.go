package main

import (
	"fmt"
)

func reverseWords(s string) string {
	b := []byte(s)
	slow, fast := 0, 0

	// 去掉首部空字符
	for fast < len(b) && b[fast] == ' ' {
		fast++
	}

	// 去掉单词间多余空字符
	for ; fast < len(b); fast++ {
		if fast - 1 > 0 && b[fast-1] == b[fast] && b[fast] == ' ' {
			continue
		}
		b[slow] = b[fast]
		slow++
	}

	// 去掉尾部空字符
	if slow - 1 > 0 && b[slow-1] == ' ' {
		b = b[:slow-1]
	} else {
		b = b[:slow]
	}

	// 翻转整个字符串
	reverseByte(&b, 0, len(b)-1)

	// 翻转单词
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverseByte(&b, i, j - 1)
		i = j + 1
	}
	return string(b)
}

func reverseByte(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}

func main() {
	s := " a good    example "
	res := reverseWords(s)
	fmt.Println(res)
}
