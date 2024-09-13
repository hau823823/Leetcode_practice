package groupAnagrams

import (
	"strings"
	"sort"
)

func GroupAnagrams(strs []string) [][]string {
    res := make([][]string, 0)
	hashMap := make(map[string][]string, 0)

	for _, str := range strs {
		org := str
		ss := strings.Split(str, "")
		sort.Strings(ss)
		str = strings.Join(ss, "")
		hashMap[str] = append(hashMap[str], org)
	}
	
	for _, v := range hashMap {
		res = append(res, v)
	}

	return res
}

// solution 2
type Key [26]byte

func strKey(str string) (key Key) {
	for i := range str {
		key[str[i]-'a']++
	}
	return
}

func GroupAnagramsOp(strs []string) [][]string {
	groups := make(map[Key][]string)
	for _, v := range strs {
		key := strKey(v)
		groups[key] = append(groups[key], v)
	}
	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}