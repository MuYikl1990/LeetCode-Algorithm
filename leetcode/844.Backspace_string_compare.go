package main

import "fmt"

func backspaceCompare(s string, t string) bool {
	sp, tp := len(s) - 1, len(t) - 1
	sCnt, tCnt := 0, 0

	for sp >= 0 || tp >= 0 {
		for sp >= 0 {
			if s[sp] == '#' {
				sCnt++
				sp--
			} else if sCnt > 0 {
				sCnt--
				sp--
			} else {
				break
			}
		}

		for tp >= 0 {
			if t[tp] == '#' {
				tCnt++
				tp--
			} else if tCnt > 0 {
				tCnt--
				tp--
			} else {
				break
			}
		}

		if sp >= 0 && tp >= 0 {
			if s[sp] != t[tp] {
				return false
			}
		} else if sp >= 0 || tp >= 0 {
			return false
		}

		sp--
		tp--
	}
	return true
}

func main() {
	s, t := "ab##", "c#d#"
	res := backspaceCompare(s, t)
	fmt.Println(res)
}
