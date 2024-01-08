package greedy_maxProfit

func maxProfit(prices []int) int {
	diff := make([]int, len(prices)-1)
	profit := 0
	for i := 1; i < len(prices); i++ {
		diff[i-1] = prices[i] - prices[i-1]
		if diff[i-1] > 0 {
			profit += diff[i-1]
		}
	}
	return profit
}
