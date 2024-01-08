package main

func permuteUnique2(nums []int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)

	backtracking(&ret, &path, nums)

	return ret
}

func backtracking(ret *[][]int, path *[]int, nums []int) {
	if len(nums) == 0 {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*ret = append(*ret, tmp)
		return
	}
	table := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := table[nums[i]]; ok {
			continue
		}

		*path = append(*path, nums[i])
		table[nums[i]] = struct{}{}

		temp := make([]int, 0)
		temp = append(temp, nums[:i]...)
		temp = append(temp, nums[i+1:]...)

		backtracking(ret, path, temp)

		*path = (*path)[:len(*path)-1]
	}
}
