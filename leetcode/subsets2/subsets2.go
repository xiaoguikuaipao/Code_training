package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0}
	fmt.Println(subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	path := []int{}
	res := make([][]int, 0, 1000)
	used := make([]bool, len(nums))
	dfs(nums, 0, used, path, &res)
	return res
}

func dfs(nums []int, startIndex int, used []bool, path []int, res *[][]int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	*res = append(*res, tmp)
	for i := startIndex; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		dfs(nums, i+1, used, path, res)
		used[i] = false
		path = path[:len(path)-1]
	}
}
