package validAnagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap := make(map[rune]int)

	for _, v := range s {
		charMap[v]++
	}

	for _, v := range t {
		c, isDup := charMap[v];
		if !isDup || c < 1 {
			return false
		}
		charMap[v]--
	}

	return true
}