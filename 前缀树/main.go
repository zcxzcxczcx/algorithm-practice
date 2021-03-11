package main

import "fmt"

func main() {
	root := Constructor()
	(&root).Insert("zcx")
	fmt.Println("rootrootroot=%v", root)
}

type Trie struct {
	Links [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		Links: [26]*Trie{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for i := 0; i < len(word); i++ {
		if root.Links[word[i]-'a'] == nil {
			root.Links[word[i]-'a'] = &Trie{}
		}
		root = root.Links[word[i]-'a']
	}
	root.isEnd = true

}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this
	for i := 0; i < len(word); i++ {
		if root.Links[word[i]-'a'] == nil {
			return false
		}
		root = root.Links[word[i]-'a']
	}
	if root.isEnd == false {
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for i := 0; i < len(prefix); i++ {
		if root.Links[prefix[i]-'a'] == nil {
			return false
		}
		root = root.Links[prefix[i]-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
