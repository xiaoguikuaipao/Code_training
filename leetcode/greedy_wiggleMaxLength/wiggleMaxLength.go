package greedy_wiggleMaxLength

func wiggleMaxLength(nums []int) int {
	differ := make([]int, 0)
	sig := 0
	for i := 1; i < len(nums); i++ {
		cur := 0
		if nums[i]-nums[i-1] < 0 {
			cur = -1
		} else if nums[i]-nums[i-1] > 0 {
			cur = 1
		} else if nums[i]-nums[i-1] == 0 {
			cur = sig
		}
		if cur != sig {
			differ = append(differ, nums[i]-nums[i-1])
			sig = cur
		} else {
			if len(differ) > 0 {
				differ[len(differ)-1] += nums[i] - nums[i-1]
			}
		}
	}
	return len(differ) + 1
}
