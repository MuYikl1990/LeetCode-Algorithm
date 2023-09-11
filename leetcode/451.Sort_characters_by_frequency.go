package main

import "fmt"

func frequencySort(s string) string {
	sArr := []byte(s)
	sMap, freqMax := make(map[byte]int), 0
	res := make([]byte, 0, len(s))
	for i := range sArr {
		sMap[sArr[i]]++
		freqMax = max451(freqMax, sMap[sArr[i]])
	}

	buckets := make([][]byte, freqMax + 1)
	for key, val := range sMap {
		buckets[val] = append(buckets[val], key)
	}

	for i := freqMax; i > 0; i-- {
		for _, ch := range buckets[i] {
			for k := i; k > 0; k-- {
				res = append(res, ch)
			}
		}
	}
	return string(res)
}

func max451(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "tree"
	res := frequencySort(s)
	fmt.Println(res)
}
