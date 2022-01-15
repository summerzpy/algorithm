package ac063

import "jz_offer/util"

//贪心策略

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profit := 0
	minValue := prices[0]
	for i := 1; i < len(prices); i++ {
		minValue = util.Min(minValue, prices[i])
		profit = util.Max(profit, prices[i]-minValue)
	}
	return profit
}
