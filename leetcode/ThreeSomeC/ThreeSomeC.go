package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSomeClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	gap := math.MaxInt
	res := 0
	seen := make(map[int]struct{})
	for i := 0; i < n-2; i++ {
		if _, ok := seen[i]; ok {
			continue
		} else {
			seen[i] = struct{}{}
		}
		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if abs(sum-target) < gap {
				res, gap = sum, abs(sum-target)
			}
			if sum == target {
				return res
			}
			if sum > target {
				k--
			} else if sum < target {
				j++
			}
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func main() {
	nums := []int{0, 0, 0}
	fmt.Println(threeSomeClosest(nums, 1))
}
