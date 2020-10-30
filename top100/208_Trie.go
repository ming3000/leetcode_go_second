package middle

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func TrieConstructor() Trie {
	return Trie{
		next:  [26]*Trie{},
		isEnd: false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		chCode := word[i] - 'a'
		if this.next[chCode] == nil {
			this.next[chCode] = new(Trie)
		} // if>>
		this = this.next[chCode]
	} // for>
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		chCode := word[i] - 'a'
		if this.next[chCode] == nil {
			return false
		} // if>>
		this = this.next[chCode]
	} // for>
	return this.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		chCode := prefix[i] - 'a'
		if this.next[chCode] == nil {
			return false
		} // if>>
		this = this.next[chCode]
	} // for>
	return true
}
