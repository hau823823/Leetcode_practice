package minCostClimbingStairs

func MinCostClimbingStairs(cost []int) int {
	l := len(cost)
	dp := make([]int, l+2)

	for i := l - 1; i >= 0; i-- {
		dp[i] = cost[i] + min(dp[i+1], dp[i+2])

	}

	return min(dp[0], dp[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
