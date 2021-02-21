package main

import "fmt"

func main() {
	t := Constructor()
	t.Insert("apple")
	fmt.Println(t.Search("apple"))
	t.Insert("app")
	fmt.Println(t.Search("app"))
	fmt.Println(t.StartsWith("appl"))
}

type Trie struct {
	Childs [26]*Trie
	isEnd  bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		Childs: [26]*Trie{},
		isEnd:  false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, c := range word {
		if this.Childs[c-'a'] == nil {
			this.Childs[c-'a'] = &Trie{Childs: [26]*Trie{}, isEnd: false}
		}
		this = this.Childs[c-'a']
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, c := range word {
		if this.Childs[c-'a'] == nil {
			return false
		}
		this = this.Childs[c-'a']
	}
	if this.isEnd == true {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, c := range prefix {
		if this.Childs[c-'a'] == nil {
			return false
		}
		this = this.Childs[c-'a']
	}
	return true
}
