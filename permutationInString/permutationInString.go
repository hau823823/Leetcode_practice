package permutationInString

func CheckInclusion(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}

	s1Count, s2Count := [26]int{}, [26]int{}
	for i := range s1 {
		s1Count[s1[i] - 'a']++
		s2Count[s2[i]-'a']++
	}

	match := make(map[[26]int]bool)
	match[s1Count] = true

	// 检查第一个窗口
	if _, exists := match[s2Count]; exists {
		return true
	}

	for i := 1; i <= l2-l1; i++ {
		// 移动窗口，增加新字符，减少旧字符
		s2Count[s2[i+l1-1]-'a']++
		s2Count[s2[i-1]-'a']--

		if _, exists := match[s2Count]; exists {
			return true
		}
	}

    return false
}

// optimize
func CheckInclusionOP(s1 string, s2 string) bool {
	l, count := 0, [26]int{}
	for i := range s1 {
		count[s1[i] - 'a']++
	}

	for j := range s2 {
		count[s2[j] - 'a']--
		if count == [26]int{} {
			return true
		}

		if j + 1 >= len(s1){
			count[s2[l] - 'a']++
			l++
		}
	}

	return false
}