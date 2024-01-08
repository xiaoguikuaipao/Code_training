package DP_uniquePaths2

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for x, flag := 0, false; x < m; x++ {
		if obstacleGrid[x][0] == 1 {
			flag = true
		}
		if flag {
			dp[x][0] = 0
		} else {
			dp[x][0] = 1
		}
	}
	for y, flag := 0, false; y < n; y++ {
		if obstacleGrid[0][y] == 1 {
			flag = true
		}
		if flag {
			dp[0][y] = 0
		} else {
			dp[0][y] = 1
		}
	}

	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			if obstacleGrid[x][y] == 1 {
				dp[x][y] = 0
			} else {
				dp[x][y] = dp[x-1][y] + dp[x][y-1]
			}
		}
	}
	return dp[m-1][n-1]
}
