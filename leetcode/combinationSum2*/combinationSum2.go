package combinationSum2

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	used := make([]int, len(candidates))
	curSum := 0
	sort.Ints(candidates)
	backtracking(0, curSum, &path, &ret, target, &used, candidates)
	return ret
}

func backtracking(start int, sum int, path *[]int, ret *[][]int, target int, used *[]int, candidates []int) {
	if sum > target {
		return
	} else if sum == target {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*ret = append(*ret, tmp)
		return
	}

	for i, e := range candidates {
		if i < start {
			continue
		}
		if i > 0 && e == candidates[i-1] && (*used)[i-1] == 0 {
			continue
		}
		sum += e
		*path = append(*path, e)
		(*used)[i] = 1
		backtracking(i+1, sum, path, ret, target, used, candidates)
		(*used)[i] = 0
		*path = (*path)[:len(*path)-1]
		sum -= e
	}
}
