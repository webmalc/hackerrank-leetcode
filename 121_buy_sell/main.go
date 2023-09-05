package buysell

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func MaxProfit(prices []int) int {
	buy := prices[0]
	profit := 0
	for _, v := range prices {
		if v < buy {
			buy = v
		} else if v-buy > profit {
			profit = v - buy
		}
	}
	return profit
}
