package main

import "fmt"

func main() {
	fmt.Println(numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
}

func numSubarraysWithSum(A []int, S int) int {
	prefixSum := 0
	table := make([]int, len(A)+1)
	count := 0
	table[0] = 1
	for _, v := range A {
		prefixSum += v
		if prefixSum >= S {
			// 每轮循环固定了窗口的右端点，找符合前缀和的左端点个数，即是符合条件的区间个数
			count += table[prefixSum-S]
		}
		table[prefixSum]++
	}
	return count
}
