package wordPattern

import "strings"

func WordPattern(pattern string, s string) bool {
	b := []byte(pattern)
	ss := strings.Split(s, " ")

	if len(b) != len(ss) {
		return false
	}

	set := make(map[byte]string) // 檢查 pattern 是否有對應的 s
	dup := make(map[string]int)  // 檢查 pattern 是否存到重複字元

	for i, v := range b {
		tmp, exists := set[v]
		dup[ss[i]]++
		if !exists{
			if dup[ss[i]] > 1 {
				return false
			}
			set[v] = ss[i]
		} else {
			if tmp != ss[i]{
				return false
			}
		}
	}

    return true
}