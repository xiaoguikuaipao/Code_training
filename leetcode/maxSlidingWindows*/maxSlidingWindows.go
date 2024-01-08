package maxSlidingWindows

func maxSlidingWindow(nums []int, k int) []int {
	right := k - 1
	left := 0
	list := make([]int, 0)
	index := 0
	ret := make([]int, len(nums)-k+1)
	for i, e := range nums {
		if i > right {
			right++
			left++
			index++
		}
		if len(list) == 0 || e < nums[list[len(list)-1]] {
			list = append(list, i)
		} else if e >= nums[list[len(list)-1]] {
			for len(list) > 0 && e >= nums[list[len(list)-1]] {
				list = list[:len(list)-1]
			}
			list = append(list, i)
		}
		if len(list) > 0 {
			for len(list) > 0 && list[0] < left {
				list = list[1:]
			}
			ret[index] = nums[list[0]]
		}
	}
	return ret
}
