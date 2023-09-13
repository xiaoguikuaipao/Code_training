package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}

var res [][]int

func permuteUnique(nums []int) [][]int {
	res = make([][]int, 0, 500)
	n := len(nums)
	sort.Ints(nums)
	used := make([]int, n)
	for i := 0; i < n; i++ {
		for i > 0 && i < n && nums[i] == nums[i-1] && used[i-1] == 0 {
			i++
		}
		if i < n {
			used[i] = 1
			path := make([]int, 0, n)
			dfs(nums, used, path, i)
			used[i] = 0
		}
	}
	return res
}

func dfs(nums, used, path []int, index int) {
	path = append(path, nums[index])
	n := len(nums)
	if len(path) == n {
		res = append(res, path)
		return
	}
	for i := 0; i < n; i++ {
		for i > 0 && i < n && nums[i] == nums[i-1] && used[i-1] == 0 {
			i++
		}
		if i < n && used[i] == 0 {
			used[i] = 1
			npath := make([]int, len(path))
			copy(npath, path)
			dfs(nums, used, npath, i)
			used[i] = 0
		}
	}
}
