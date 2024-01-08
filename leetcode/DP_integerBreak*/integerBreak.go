package DP_integerBreak

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = getMax(j*(i-j), j*dp[i-j], dp[i])
		}
	}
	return dp[n]
}

func getMax(x int, y int, z int) int {
	max := x
	if y > max {
		max = y
	}
	if z > max {
		max = z
	}
	return max
}
