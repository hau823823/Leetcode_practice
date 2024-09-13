package bestTimeBuySellStock1

func MaxProfitTraverse(prices []int) int {

	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

    return maxProfit
}

func MaxProfitGreedyApproach(prices []int) int {

	// find the lowest buy price
	minBuyPrice := prices[0]

	// check the maxProfit
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minBuyPrice {
			minBuyPrice = prices[i]
		} else if (prices[i] - minBuyPrice) > maxProfit {
			maxProfit = prices[i] - minBuyPrice
		}
	}

	return maxProfit
}

func MaxProfitDP(prices []int) int {

	n := len(prices)
	if n == 0 {
		return 0
	}
	
	// dp 陣列的維度和意義：
	// dp 是一個 2D 陣列，其中：
	
	// 第一個維度的大小為 n（價格陣列的長度），代表每一天。
	// 第二個維度有兩個元素，分別是 0 和 1。
	// dp[i][0]：表示在第 i 天持有股票時能獲得的最大利潤。
	// dp[i][1]：表示在第 i 天不持有股票時能獲得的最大利潤。

	dp := make([][2]int, n)
	dp[0][0] = -prices[0] // 第一天持有股票的話會直接虧
	dp[0][1] = 0

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] + prices[i])
	}

	return max(dp[n-1][0], dp[n-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}