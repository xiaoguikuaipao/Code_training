package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-4, -1, -1, 0, 1, 2}))
}

// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

// 注意：答案中不可以包含重复的三元组。
// 示例 1：

// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 解释：
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
// 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
// 注意，输出的顺序和三元组的顺序并不重要。

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			}
		}
	}
	return result
}
