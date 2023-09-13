package ThreeSum

import "sort"

func threesum(nums []int) [][]int {
	sort.Ints(nums)
	result, start, end, index := make([][]int, 0), 0, 0, 0
	Sum := 0
	length := len(nums)
	for index = 0; index < length-1; index++ {
		if nums[index] > 0 {
			break
		}
		if index >= 1 && nums[index] == nums[index-1] {
			continue
		}
		start = index + 1
		end = length - 1
		for start < length && start < end {
			Sum = nums[start] + nums[end] + nums[index]
			if Sum == 0 {
				result = append(result, []int{nums[start], nums[end], nums[index]})
				start++
				for nums[start] == nums[start-1] && start < end {
					start++
				}
			} else if Sum < 0 {
				start++
				for nums[start] == nums[start-1] && start < end {
					start++
				}
			} else if Sum > 0 {
				end--
				for nums[end] == nums[end+1] && start < end {
					end--
				}
			}
		}
	}
	return result
}
