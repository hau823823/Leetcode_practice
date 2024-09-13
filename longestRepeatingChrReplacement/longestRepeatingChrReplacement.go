package longestRepeatingChrReplacement

func CharacterReplacement(s string, k int) int {
	count := make(map[byte]int)
	res, max, left := 0, 0, 0

	for right := range s {
		count[s[right]] += 1
		max = maxInt(max, count[s[right]])

		if right - left + 1 - max > k {
			count[s[left]] -= 1
			left++
		}

		res = maxInt(res, right - left + 1)
	}

	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}