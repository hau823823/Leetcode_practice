package implementStr

import (
	"strings"
)

// 第一次自己構建
// 問題在於可能會有重複的字元覆蓋指引位置
func StrStr(haystack string, needle string) int {
	hayMap := make(map[byte]int)

	hay := []byte(haystack)
	for i, str := range hay {
		hayMap[str] = i
	}

	nee := []byte(needle)
	idx, ok := hayMap[nee[0]]
	index := 0
	if ok {
		index += idx
	} else {
		return -1
	}

	check := 1
	for i, str := range nee {
		idx, ok := hayMap[str]
		if ok && (idx - index) == i {
			continue
		} else {
			check = check * -1
			break
		}

	}

    if check == 1 {
		return index
	}
	return -1
}

func StrStrUseStrings(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func StrStrKMP(haystack string, needle string) int {
	return -1
}