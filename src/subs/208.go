package subs

type Trie struct {
	words map[string]bool
	prefix map[string]bool
}


/** Initialize your data structure here. */
func Constructor208() Trie {
	tire := Trie{}
	tire.words = make(map[string]bool)
	tire.prefix = make(map[string]bool)
	return tire
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	this.words[word] = true
	this.prefix[word] = true
	for i:=0;i<len(word);i++{
		this.prefix[word[0:i]] = true
	}
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.words[word]
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.prefix[prefix]
}