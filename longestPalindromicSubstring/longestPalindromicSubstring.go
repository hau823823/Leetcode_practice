package longestPalindromicSubstring

import "bytes"

// Brute Force
func LongestPalindromeBF(s string) string {

	tmp := ""
	l := len(s)

	if l == 1 {
		return s
	}

	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if CheckPalindrome(s[i:j+1]) && len(s[i:j+1]) > len(tmp) {
				tmp = s[i : j+1]
			}
		}

	}

	return tmp
}

func CheckPalindrome(s string) bool {
	l := len(s)
	left, right := 0, 0
	if l%2 == 1 {
		left = l / 2
		right = left + 2
	} else {
		left = l / 2
		right = left + 1
	}

	b := []byte(s)
	for i := 0; i < left; i++ {
		if b[left-i-1] != b[right+i-1] {
			return false
		}
	}
	return true
}

// Dynamic Programming 1
func LongestPalindromeDP(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	start, end := 0, 0
	for j := 1; j < len(s); j++ {
		for i := j - 1; i >= 0; i-- {
			if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j-i+1 > end-start+1 {
					start = i
					end = j
				}
			}
		}
	}

	return s[start : end+1]
}

// Dynamic Programming 2
func LongestPalindromeDP2(s string) string {
	if len(s) == 0 {
		return ""
	}

	n := len(s)
	left, lenMax := 0, 1
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			dp[j][i] = (s[i] == s[j]) && (i-j < 2 || dp[j+1][i-1])
			if dp[j][i] && lenMax < i-j+1 {
				lenMax = i - j + 1
				left = j
			}
		}
	}
	return s[left : left+lenMax]
}

// Manacher’s Algorithm
func LongestPalindromeManachers(s string) string {
	// Step 1
	// 為了使奇偶數長的 s 字串都可以繼續進行，因此加入虛假字元填充空格
	// 这样做的原因是为了统一处理原始字符串中的奇数长度和偶数长度的回文子串。
	// 通过这种插入，所有可能的回文子串在新字符串中都将是奇数长度，这样就可以统一处理。
	var buffer bytes.Buffer
	buffer.WriteByte('*')
	for i := 0; i < len(s); i++ {
		buffer.WriteByte(s[i])
		buffer.WriteByte('*')
	}
	sp := buffer.String()

	// Step2
	dp := make([]int, len(sp))
	center := 0
	radius := 0
	for center < len(sp) {
		for center-radius-1 >= 0 && center+radius+1 < len(sp) {
			if sp[center-radius-1] == sp[center+radius+1] {
				radius++
			} else {
				break
			}
		}
		dp[center] = radius

		oldCenter, oldRadius := center, radius
		upperBound := oldCenter + oldRadius

		center++
		radius = 0
		for center <= upperBound {
			mirror := 2*oldCenter - center
			mmr := upperBound - center

			if dp[mirror] == mmr {
				radius = mmr
				break
			}

			dp[center] = minOf(dp[mirror], mmr)
			center++
		}
	}

	//Step 3
	max := 0
	maxIdx := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
			maxIdx = i
		}
	}
	ss := sp[maxIdx-max : maxIdx+max+1]
	buffer.Reset()
	for i := 1; i < len(ss); i += 2 {
		buffer.WriteByte(ss[i])
	}
	return buffer.String()
}

func minOf(a, b int) int {
	if a < b {
		return a
	}
	return b
}
