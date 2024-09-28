package bestTimeBuySellStocksWithCooldown

import "math"

func MaxProfit(prices []int) int {
	
	// sold: 表示當天賣出後的最大利潤
	// hold: 表示當天持有股票後的最大利潤，初始狀態未持有股票故使用極小值
	// rest: Cool down 不進行任何操作後的最大利潤
	sold, hold, rest := 0, math.MinInt, 0

	for _, price := range prices {
		prevSold := sold
		sold = hold + price
		hold = max(hold, rest - price)
		rest = max(rest, prevSold)

	}

	return max(sold, rest)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}