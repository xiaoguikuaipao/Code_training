package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

var res [][]int

func combine(n, k int) [][]int {
	res = make([][]int, 0, 500)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	path := make([]int, 0, n)
	for start := 0; start < n; start++ {
		path = append(path, nums[start])
		dfs(nums, n, k, path, start)
		path = path[:len(path)-1]
	}
	return res
}

func dfs(nums []int, n, k int, path []int, start int) {
	if n-start < k-len(path) {
		return
	}
	if len(path) == k {
		store := make([]int, k)
		copy(store, path)
		res = append(res, store)
	}
	for i := start + 1; i < n; i++ {
		path = append(path, nums[i])
		dfs(nums, n, k, path, i)
		path = path[:len(path)-1]
	}
}
