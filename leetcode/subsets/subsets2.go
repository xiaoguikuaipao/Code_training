package main

func subsets2(nums []int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	ret = append(ret, []int{})
	backtracking(&ret, &path, 0, nums)

	return ret
}

func backtracking(ret *[][]int, path *[]int, start int, nums []int) {
	for i := start; i < len(nums); i++ {
		*path = append(*path, nums[i])
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*ret = append(*ret, tmp)
		backtracking(ret, path, i+1, nums)
		*path = (*path)[:len(*path)-1]
	}
}
