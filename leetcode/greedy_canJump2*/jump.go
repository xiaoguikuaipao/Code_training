package greedy_canJump2

func jump(nums []int) int {
	firstAchieve := make([]int, len(nums))
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i+nums[i] > farthest {
			for j := farthest + 1; j <= i+nums[i] && j < len(nums); j++ {
				firstAchieve[j] = i
			}
			farthest = i + nums[i]
			if farthest >= len(nums)-1 {
				return getJump(firstAchieve, len(nums)-1)
			}
		}
	}
	return 0
}

func getJump(achieve []int, i int) int {
	cnt := 0
	for i != 0 {
		i = achieve[i]
		cnt++
	}
	return cnt
}
