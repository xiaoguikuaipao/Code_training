package main

import "fmt"

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。
示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3

	1 [3  -1  -3] 5  3  6  7       3
	1  3 [-1  -3  5] 3  6  7       5
	1  3  -1 [-3  5  3] 6  7       5
	1  3  -1  -3 [5  3  6] 7       6
	1  3  -1  -3  5 [3  6  7]      7
*/

func maxSlidingWindow(nums []int, k int) []int {
	st := make([]int, 0, k)
	start := 0
	end := k - 1
	result := make([]int, 0, k)
	for i, v := range nums {
		for len(st) > 0 {
			top := st[len(st)-1]
			if v < nums[top] && top >= start {
				break
			}
			st = st[:len(st)-1]
		}
		for len(st) > 0 && st[0] < start {
			st = st[1:]
		}
		st = append(st, i)
		st = append(st, i)
		if i == end {
			start++
			end++
			result = append(result, nums[st[0]])
		}
	}
	return result
}

func main() {
	nums := []int{1, 3, 1, 2, 0, 5}
	fmt.Println(maxSlidingWindow(nums, 3))
}
