package longestPalindrome

func LongestPalindrome(s string) int {
	res, tmp := 0, 0
	m := make(map[rune]int)
	for _, b := range s {
		m[b]++

		if m[b] % 2 == 0 {
			res += 2
			tmp--
		} else {
			tmp++
		}
	}

	if tmp > 0 {
		res++
	}

	return res
}