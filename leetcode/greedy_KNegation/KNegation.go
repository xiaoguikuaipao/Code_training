package greedy_KNegation

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if k == 0 {
			return getSum(nums)
		}
		if nums[i] < 0 {
			nums[i] = -nums[i]
			k--
			if i == len(nums)-1 && k > 0 {
				i = 0
			}
			continue
		}
		if nums[i] > 0 {
			sort.Ints(nums)
			nums[0] = int(math.Pow(-1, float64(k))) * nums[0]
			return getSum(nums)
		}
	}
	return 0
}

func getSum(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	return sum
}
