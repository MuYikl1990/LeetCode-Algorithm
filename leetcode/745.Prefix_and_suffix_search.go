package main

import "fmt"

type TrieNode struct {
	strs     []int
	children [26]*TrieNode
}

type WordFilter struct {
	prefix *TrieNode
	suffix *TrieNode
}

func Constructor745(words []string) WordFilter {
	prefix, suffix := build(words, true), build(words, false)
	return WordFilter{prefix, suffix}
}

func build(words []string, isPre bool) *TrieNode {
	root := &TrieNode{}
	for pos, word := range words {
		tmp := root
		i := 0
		if !isPre {
			i = len(word) - 1
		}
		for i >= 0 && i < len(word) {
			if tmp.children[word[i] - 'a'] == nil {
				tmp.children[word[i] - 'a'] = &TrieNode{}
			}
			tmp = tmp.children[word[i] - 'a']
			tmp.strs = append(tmp.strs, pos)
			if isPre {
				i++
			} else {
				i--
			}
		}
	}
	return root
}

func (this *WordFilter) F(pref string, suff string) int {
	preList := query(this.prefix, pref, true)
	sufList := query(this.suffix, suff, false)
	i, j := len(preList) - 1, len(sufList) - 1
	for i >= 0 && j >= 0 {
		if preList[i] == sufList[j] {
			return preList[i]
		} else if preList[i] > sufList[j] {
			i--
		} else {
			j--
		}
	}
	return -1
}

func query(trie *TrieNode, str string, isPre bool) []int {
	i := 0
	if !isPre {
		i = len(str) - 1
	}

	for i >= 0 && i < len(str) {
		if trie.children[str[i] - 'a'] == nil {
			return []int{}
		}
		trie = trie.children[str[i] - 'a']
		if isPre {
			i++
		} else {
			i--
		}
	}
	return trie.strs
}

func main() {
	words, pref, suff := []string{"apple", "address", "allocate"}, "a", "e"
	obj := Constructor745(words)
	param := obj.F(pref,suff)
	fmt.Println(param)
}
