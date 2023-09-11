package main

import "fmt"

func isNumber(s string) bool {
	// 数字标识，点标识，e标识
	numFlag, dotFlag, eFlag := false, false, false
	for i := 0; i < len(s); i++ {
		// 数字标识为true
		if s[i] >= '0' && s[i] <= '9' {
			numFlag = true
				// s[i]为'.'，需要点未出现过且e未出现过
		} else if s[i] == '.' && !dotFlag && !eFlag {
			dotFlag = true
			   // s[i]为e，需要e未出现过且之前有数字
		} else if (s[i] == 'e' || s[i] == 'E') && !eFlag && numFlag {
			eFlag = true
			numFlag = false  // 出现e之后重新判断整数
			   // s[i]为+-，只能出现在首位或者e之后
		} else if (s[i] == '+' || s[i] == '-') && (i == 0 || s[i - 1] == 'e' || s[i - 1] == 'E') {
			continue
		} else {
			// 其他情况非法
			return false
		}
	}
	// 根据 numFlag 进行判断，避免'.' 或 '1e'这种测试样例，保证出现数字
	return numFlag
}

func main() {
	s := "e"
	res := isNumber(s)
	fmt.Println(res)
}
