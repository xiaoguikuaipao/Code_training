package dp_targetSum

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	sum := 0
	for _, e := range nums {
		sum += e
	}
	if (target+sum)%2 != 0 || target+sum < 0 {
		return 0
	}
	positiveSum := (target + sum) / 2
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, positiveSum+1)
	}

	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j <= positiveSum; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j-nums[i-1]] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][positiveSum]
}
