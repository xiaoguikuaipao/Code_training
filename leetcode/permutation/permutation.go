package main

import "fmt"

func main() {
	nums := []int{0, 1}
	fmt.Println(permute(nums))
}

var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	n := len(nums)
	used := make([]int, n)
	for i, _ := range nums {
		used[i] = 1
		path := make([]int, 0, n)
		dfs(i, nums, used, path)
		used[i] = 0
	}
	return res
}

func dfs(index int, nums []int, used []int, path []int) {
	path = append(path, nums[index])
	n := len(nums)
	if len(path) == n {
		res = append(res, path)
		return
	}
	for i := 0; i < n; i++ {
		if used[i] == 0 {
			used[i] = 1
			newpath := make([]int, len(path))
			copy(newpath, path)
			dfs(i, nums, used, newpath)
			used[i] = 0
		}
	}
}
