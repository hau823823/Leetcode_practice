package designAddAndSearchWordsDataStructure

type WordDictionary struct {
	children map[rune]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{make(map[rune]*WordDictionary), false}
}

func (wd *WordDictionary) AddWord(word string) {
	cur := wd
	for _, c := range word {
		if _, exists := cur.children[c]; !exists {
			cur.children[c] = &WordDictionary{make(map[rune]*WordDictionary), false}
		}
		cur = cur.children[c]
	}
	cur.isEnd = true
}

func (wd *WordDictionary) Search(word string) bool {
	cur := wd
	for i, c := range word {
		if _, exists := cur.children[c]; !exists {
			if c != '.' {
				return false
			}
			for _, child := range cur.children {
				if child.Search(word[i + 1:]) {
                    return true
                }
			}

			return false
		}
		cur = cur.children[c]
	}

	return cur.isEnd
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
