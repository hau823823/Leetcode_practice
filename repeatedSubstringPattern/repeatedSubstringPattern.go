package repeatedSubstringPattern

import "strings"

// Brute Force
func RepeatedSubstringPattern(s string) bool {
	result := false

	for i := 0; i < len(s)/2; i++ {
		tmp := s[:i+1]

		if len(s)%len(tmp) == 0 {
			j := i + 1
			for j+len(tmp) < len(s)+1 {
				if tmp == s[j:j+len(tmp)] {
					result = true
				} else {
					result = false
					break
				}
				j += len(tmp)
			}
		}

		if result {
			return result
		}
	}

	return result
}

// Rotation
// time complexity: O(n^2)
func RepeatedSubstringPatternRt(s string) bool {
	l := len(s)

	for i := 1; i < l; i++ {
		tmp := s[l-i:] + s[:l-i] // O(n)
		if s == tmp {
			return true
		}
	}

	return false
}

// Concatenation
func RepeatedSubstringPatternCt(s string) bool {
	doubled := s + s
	return strings.Contains(doubled[1:len(doubled)-1], s)
}

// KMP
func RepeatedSubstringPatternKMP(s string) bool {
	n := len(s)
	dp := make([]int, n + 1)
	i, j := 1, 0 // i 用于遍历字符串 s 的索引，j 用于标记当前位置前的子串中，最长的相同前缀和后缀的长度

	for i < n {
		if s[i] == s[j] {
			j++
			i++
			dp[i] = j
		} else if j == 0 {
			i++
		} else {
			j = dp[j]
		}
	}

	return dp[n] != 0 && n % (n - dp[n]) == 0
}
