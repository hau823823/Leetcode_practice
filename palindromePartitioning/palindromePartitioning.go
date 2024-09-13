package palindromePartitioning

func Partition(s string) [][]string {
	res := [][]string{}
	cur := []string{}
	l := len(s)

	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == l {
			cpy := append([]string{}, cur...)
			res = append(res, cpy)
			return
		}

		for i := idx; i < l; i++ {
			sub := s[idx:i+1]
			if isPalindrome(sub) {
				cur = append(cur, sub)
				backtrack(i+1)
				cur = cur[:len(cur)-1]
			}
		}
	}

	backtrack(0)
	return res
}

func isPalindrome(s string) bool {
	left, right := 0, len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}