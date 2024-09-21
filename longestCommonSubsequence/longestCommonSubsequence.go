package longestCommonSubsequence

func LongestCommonSubsequence(text1 string, text2 string) int {

	n, m := len(text1), len(text2)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		dp[i][0] = 0
	}
	for j := 0; j < m; j++ {
		dp[0][j] =0
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i-1][j-1] + 1
		}
	}

    return 0
}