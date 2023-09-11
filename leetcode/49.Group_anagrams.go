package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)
	for _, str := range strs {
		var cnt [26]int
		for _, s := range str {
			cnt[s - 'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}

	var res [][]string
	for _, val := range mp {
		res = append(res, val)
	}
	return res
}

func main() {
	strs := []string{"eat","tea","tan","ate","nat","bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)
}
