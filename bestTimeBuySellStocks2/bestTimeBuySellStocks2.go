package bestTimeBullSellStocks2

func MaxProfit(prices []int) int{
	profit := 0
    minBuyPrice := prices[0]
    for i := 1; i < len(prices); i++ {
        if prices[i] < minBuyPrice {
            minBuyPrice = prices[i]
        } else if prices[i] - minBuyPrice > profit {
            profit = prices[i] - minBuyPrice
        }
    }
    
    return profit
}