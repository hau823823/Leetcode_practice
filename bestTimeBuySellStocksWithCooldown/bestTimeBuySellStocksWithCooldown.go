package bestTimeBuySellStocksWithCooldown

import "math"

func MaxProfit(prices []int) int {

	// sold: 表示當天賣出後的最大利潤
	// hold: 表示當天持有股票後的最大利潤，初始狀態未持有股票故使用極小值
	// rest: Cool down 不進行任何操作後的最大利潤
	sold, hold, rest := 0, math.MinInt, 0

	for _, price := range prices {
		prevSold := sold             // 保存之前一天的 sold，因為在更新 rest 會需要
		sold = hold + price          // 當天賣出股票，則當天 sold 是前一天持有股票的狀態再加上今天賣出的價格
		hold = max(hold, rest-price) // 今天如果選擇買入股票，那麼狀態 hold 就是前一天的 rest 狀態減去今天的股票價格。如果不買入股票，則保持前一天 hold 的狀態
		rest = max(rest, prevSold)   // 今天休息的狀態 rest 可以是前一天的休息狀態，也可以是前一天賣出股票後的狀態（因為股票賣出後至少需要一天的冷卻期）

	}

	return max(sold, rest)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
