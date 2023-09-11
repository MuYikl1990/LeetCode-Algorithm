package main

import "fmt"

// 36进制由0-9，a-z，共36个字符表示。
// 要求按照加法规则计算出任意两个36进制正整数的和，如1b + 2x = 48  （解释：47+105=152）
// 要求：不允许使用先将36进制数字整体转为10进制，相加后再转回为36进制的做法

func add36Strings(num1, num2 string) string {
	getInt := func(char byte) int {
		if char > '9' {
			return int(char - 'a' + 10)
		} else {
			return int(char - '0')
		}
	}
	getChar := func(num byte) byte {
		if num < 10 {
			return '0' + num
		}
		return num - 10 + 'a'
	}

	var res []byte
	i, j, carry := len(num1) - 1, len(num2) - 1, 0
	for i >= 0 || j >= 0 || carry > 0 {
		if i >= 0 {
			carry += getInt(num1[i])
			i--
		}
		if j >= 0 {
			carry += getInt(num2[j])
			j--
		}
		res = append(res, byte(carry % 36))
		carry /= 36
	}

	left, right := 0, len(res) - 1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}

	for k := range res {
		res[k] = getChar(res[k])
	}
	return string(res)
}

func main() {
	num1, num2 := "1b", "2b"
	res := add36Strings(num1, num2)
	fmt.Println(res)
}
