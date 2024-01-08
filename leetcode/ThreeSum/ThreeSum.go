package ThreeSum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	for i, e := range nums {
		if e > 0 {
			break
		}
		if i > 0 && e == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if e+(nums[left]+nums[right]) == 0 {
				ret = append(ret, []int{e, nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
			if e+(nums[left]+nums[right]) > 0 {
				right--
			}
			if e+(nums[left]+nums[right]) < 0 {
				left++
			}
		}
	}
	return ret
}
