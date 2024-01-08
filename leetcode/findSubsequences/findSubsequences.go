package main

import "fmt"

func findSubsequences(nums []int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)

	backtracking(&ret, &path, 0, nums)

	return ret
}

func backtracking(ret *[][]int, path *[]int, start int, nums []int) {
	table := make(map[int]struct{})
	for i := start; i < len(nums); i++ {
		if len(*path) == 0 || nums[i] >= (*path)[len(*path)-1] {
			if _, ok := table[nums[i]]; ok {
				continue
			}

			*path = append(*path, nums[i])
			table[nums[i]] = struct{}{}

			if len(*path) >= 2 {
				tmp := make([]int, len(*path))
				copy(tmp, *path)
				*ret = append(*ret, tmp)
			}

			backtracking(ret, path, i+1, nums)

			*path = (*path)[:len(*path)-1]
		}
	}
}

func main() {
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}
