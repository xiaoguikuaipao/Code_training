package DP_lastStones2

func lastStoneWeightII(stones []int) int {
	n := len(stones)
	sum := 0
	for _, e := range stones {
		sum += e
	}
	target := sum / 2
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = 0
	}
	for j := 0; j <= target; j++ {
		if j >= stones[0] {
			dp[0][j] = stones[0]
		} else {
			dp[0][j] = 0
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= target; j++ {
			if j >= stones[i] {
				dp[i][j] = getMax(dp[i-1][j-stones[i]]+stones[i], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return sum - 2*dp[n][target]
}

func getMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
