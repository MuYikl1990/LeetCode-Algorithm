package main

import "fmt"

type Trie struct {
	children [26]*Trie
	isEnd    bool
}


func Constructor208() Trie {
	return Trie{}
}


func (this *Trie) Insert(word string)  {
	for _, ch := range word {
		ch -= 'a'
		if this.children[ch] == nil {
			this.children[ch] = &Trie{}
		}
		this = this.children[ch]
	}
	this.isEnd = true
}


func (this *Trie) Prefix(prefix string) *Trie {
	for _, ch := range prefix {
		ch -= 'a'
		if this.children[ch] == nil {
			return nil
		}
		this = this.children[ch]
	}
	return this
}


func (this *Trie) Search(word string) bool {
	node := this.Prefix(word)
	return node != nil && node.isEnd == true
}


func (this *Trie) StartsWith(prefix string) bool {
	return this.Prefix(prefix) != nil
}

func main() {
	word := "seller"
	prefix := "sea"
	obj := Constructor208()
	obj.Insert(word)
	param1 := obj.Search("sel")
	param2 := obj.StartsWith(prefix)
	fmt.Println(param1, param2)
}
