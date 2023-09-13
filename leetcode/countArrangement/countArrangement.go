package main

import (
	"fmt"
)

func main() {
	fmt.Println(countArrangement(4))
	fmt.Println(res)
}

var res [][]int

func countArrangement(n int) int {
	res = make([][]int, 0, 100)
	nums := make([]int, n+1)
	used := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		nums[i] = i
	}
	path := make([]int, 0, n+1)
	path = append(path, 0)
	dfs(path, nums, used)
	return len(res)
}

func dfs(path []int, nums, used []int) {
	index := len(path)
	if len(nums) == index {
		tmp := make([]int, index)
		copy(tmp, path)
		res = append(res, tmp)
	}
	for _, num := range nums {
		if num != 0 && used[num] == 0 && (num%index == 0 || index%num == 0) {
			path = append(path, num)
			used[num] = 1
			dfs(path, nums, used)
			path = path[:len(path)-1]
			used[num] = 0

		}
	}
}
