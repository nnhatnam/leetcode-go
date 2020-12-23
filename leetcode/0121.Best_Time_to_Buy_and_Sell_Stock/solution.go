package _121_Best_Time_to_Buy_and_Sell_Stock

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	buyPrice := prices[0]
	sellPrice := prices[0]

	profit := sellPrice - buyPrice

	for i := 1; i < len(prices); i++ {
		if prices[i] - buyPrice > profit {
			sellPrice = prices[i]
			profit = prices[i] - buyPrice
		}
	}

	return profit
}
