package DP_findMaxForm

func findMaxForm(strs []string, m int, n int) int {
	size := len(strs)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i := 0; i < size; i++ {
		cnt0 := 0
		cnt1 := 0
		for _, e := range strs[i] {
			if e == '1' {
				cnt1++
			} else {
				cnt0++
			}
		}
		for j := m; j >= cnt0; j-- {
			for k := n; k >= cnt1; k-- {
				dp[j][k] = getMax(dp[j][k], dp[j-cnt0][k-cnt1]+1)
			}
		}
	}
	return dp[m][n]
}

func getMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
