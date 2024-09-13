package coinChange

import "math"

func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)

	for i := 1; i <= amount; i++ {
		minCoin := math.MaxInt
		for _, coin := range coins {
			if i - coin >= 0 && dp[i - coin] != -1 {
				minCoin = min(minCoin, dp[i - coin] + 1)
			}
		}
		if minCoin == math.MaxInt {
			dp[i] = -1
		} else {
			dp[i] = minCoin
		}
	}
	return dp[amount]
}

func CoinChangeOp(coins []int, amount int) int {
	limit := amount + 1
	dp := make([]int, limit + 1)

	var minCoin int
	for i := 1; i < limit; i++ {
		minCoin = limit

		for _, coin := range coins {
			if coin > i {
				continue
			}
			minCoin = min(minCoin, dp[i - coin] + 1)
		}

		dp[i] = minCoin
	}

	if dp[amount] == limit {
		return -1
	}
 
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}