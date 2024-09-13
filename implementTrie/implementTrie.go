package implementTrie

// time complexity
// Insert Operation: O(m), where m is the length of the word being inserted.
// Search Operation: O(m), where m is the length of the word being searched.
// StartsWith Operation: O(m), where m is the length of the prefix.

// Space complexity: O(N×M)
// where N is the number of words and M is the average length of the words. 
// Each node may have up to 26 children (for each letter of the alphabet in English).

type Trie struct {
    children [26]*Trie
	isEnd bool
}


func Constructor() Trie {
    return Trie{}
}


func (t *Trie) Insert(word string)  {
	cur := t
	for _, c := range word {
		idx := c - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = &Trie{}
		}
		cur = cur.children[idx]
	}
	cur.isEnd = true
}


func (t *Trie) Search(word string) bool {
    cur := t
	for _, c := range word {
		idx := c - 'a'
		if cur.children[idx] == nil {
			return false
		}
		cur = cur.children[idx]
	}
	return cur.isEnd
}


func (t *Trie) StartsWith(prefix string) bool {
    cur := t
	for _, c := range prefix {
		idx := c - 'a'
		if cur.children[idx] == nil {
			return false
		}
		cur = cur.children[idx]
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