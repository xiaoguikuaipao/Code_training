package DP_minCostsClimbingStairs

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < len(dp); i++ {
		dp[i] = getMin(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(dp)-1]
}

func getMin(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
