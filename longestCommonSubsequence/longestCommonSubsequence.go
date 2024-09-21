package longestCommonSubsequence

// 2-D dynamic
// Bottom Up
func LongestCommonSubsequence(text1 string, text2 string) int {

	n, m := len(text1), len(text2)
	dp := make([][]int, n + 1)

	for i := 0; i < n + 1; i++ {
		dp[i] = make([]int, m + 1)
	}

	for i := n-1; i >= 0; i-- {
		for j := m-1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}


    return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}