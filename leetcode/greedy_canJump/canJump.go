package greedy_canJump

func canJump(nums []int) bool {
	farthest := nums[0]
	if len(nums) == 1 {
		return true
	}
	for i := 0; i < len(nums)-1; i++ {
		if i > farthest {
			return false
		}
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if i+nums[i] >= len(nums)-1 {
			return true
		}
	}
	return false
}
