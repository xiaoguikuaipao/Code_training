package main

func canPartition2(nums []int) bool {
	target := 0
	for _, e := range nums {
		target += e
	}
	if target%2 != 0 {
		return false
	}
	target /= 2
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	for i := 0; i < len(nums); i++ {
		dp[i][0] = 0
	}
	for j := 0; j <= target; j++ {
		if j >= nums[0] {
			dp[0][j] = nums[0]
		} else {
			dp[0][j] = 0
		}
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j <= target; j++ {
			if j >= nums[i] {
				dp[i][j] = getMax(dp[i-1][j-nums[i]]+nums[i], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][target] == target
}

func getMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
