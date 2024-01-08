package greedy_maxSubArray

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	path := make([]int, 0)
	cur := 0
	for i := 0; i < len(nums); i++ {
		path = append(path, nums[i])
		cur += nums[i]
		if cur > maxSum {
			maxSum = cur
		}

		if cur <= 0 {
			path = path[:0]
			cur = 0
			continue
		}
	}
	return maxSum
}
