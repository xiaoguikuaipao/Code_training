package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, -1, -1, 1, 1, 2, 2}
	target := 0
	fmt.Println(fourSum(nums, target))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	if target > 0 && nums[0] > target {
		return [][]int{}
	}
	if target < 0 && nums[len(nums)-1] < target {
		return [][]int{}
	}
	for i, e1 := range nums {

		//pay attention to deduplication
		if i > 0 && e1 == nums[i-1] {
			continue
		}

		res := target - e1
		if res > 0 && nums[len(nums)-1]*3 < res {
			continue
		}
		if res < 0 && i+1 < len(nums) && nums[i+1]*3 > res {
			continue
		}
		for j, size := i+1, len(nums); j < size; j++ {

			//pay attention to deduplication
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := size - 1
			for left < right {
				if e1+nums[j]+nums[left]+nums[right] == target {
					ret = append(ret, []int{e1, nums[j], nums[left], nums[right]})
					left++
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					right--
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				}
				if e1+nums[j]+nums[left]+nums[right] < target {
					left++
				}
				if e1+nums[j]+nums[left]+nums[right] > target {
					right--
				}
			}
		}

	}
	return ret
}
