package palindromicSubstring

// brute forece
func CountSubstringBF(s string) int {
	result := 0
	check := func(i, j int) {
		for i >= 0 && j < len(s) && s[i] == s[j] {
			result++
			i--
			j++
		}
	}

	for i := 0; i < len(s); i++ {
		check(i, i)		// 检查以字符 s[i] 为中心的所有奇数长度的回文子串
		check(i, i+1)	// 如果 s[i] 和 s[i+1] 相同，则检查以这两个字符为中心的所有偶数长度的回文子串
	}

	return result
}

// 2-D dp
func CountSubstringDP(s string) int {
	result := 0
	dp := make([][]bool, len(s))		// dp[i][j] 表示从索引 i 到 j 的子串是否为回文

	for i := len(s) - 1; i >= 0; i-- {	// 从字符串的末尾开始向前迭代，这是因为计算 dp[i][j] 依赖于 dp[i+1][j-1]，所以需要先计算后面的值
		dp[i] = make([]bool, len(s))

		for j := i; j < len(s); j++ {	// 可以保证当检查子串 s[i] 到 s[j] 是否为回文时，所有需要的 dp 值都已经计算过了
			if s[i] == s[j] && (j - i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				result++
			}
		}
	}

	return result
}

// 
func CountSubstringMC(s string) int {
	return 0
}