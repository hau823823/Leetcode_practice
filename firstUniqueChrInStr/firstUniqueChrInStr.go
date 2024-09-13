package firstUniqueChrInStr

func FirstUniqChar(s string) int {
	charMap := make(map[rune]int)
	for _, v := range s {
		charMap[v]++
	}
	for i, v := range s {
		if charMap[v] == 1 {
			return i
		}
	}
	return -1
}