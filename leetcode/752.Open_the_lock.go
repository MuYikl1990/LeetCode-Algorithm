package main

import "fmt"

func openLock(deadends []string, target string) int {
	deadMap := make(map[string]bool)
	for _, dead := range deadends {
		deadMap[dead] = true
	}

	if deadMap["0000"] {
		return -1
	}

	if target == "0000" {
		return 0
	}

	var queue1, queue2 []string
	map1, map2 := make(map[string]int), make(map[string]int)
	queue1 = append(queue1, "0000")
	queue2 = append(queue2, target)
	map1["0000"] = 0
	map2[target] = 0

	res := -1
	for len(queue1) != 0 && len(queue2) != 0 {
		if len(queue1) <= len(queue2) {
			res = bfs(&queue1, map1, map2, deadMap)
		} else {
			res = bfs(&queue2, map2, map1, deadMap)
		}
		if res != -1 {
			return res
		}
	}
	return res
}

func bfs(queue *[]string, map1, map2 map[string]int, deadMap map[string]bool) int {
	n := len(*queue)
	for i := 0; i < n; i++ {
		cur := []byte((*queue)[i])
		step := map1[(*queue)[i]]
		for j := 0; j < 4; j++ {
			origin := cur[j]
			for k := -1; k <= 1; k += 2 {
				if k == -1 {
					cur[j] = pre(cur[j])
				} else {
					cur[j] = next(cur[j])
				}
				if !deadMap[string(cur)] && map1[string(cur)] == 0 {
					if val, ok := map2[string(cur)]; ok {
						return step + 1 + val
					}
					fmt.Println(string(cur))
					map1[string(cur)] = step + 1
					*queue = append(*queue, string(cur))
				}
				cur[j] = origin
			}
		}
	}
	*queue = (*queue)[n:]
	return -1
}

func next(b byte) byte {
	if b == '9' {
		return '0'
	}
	return b + 1
}

func pre(b byte) byte {
	if b == '0' {
		return '9'
	}
	return b - 1
}

func main() {
	deadends, target := []string{"8888"}, "0009"
	res := openLock(deadends, target)
	fmt.Println(res)
}
