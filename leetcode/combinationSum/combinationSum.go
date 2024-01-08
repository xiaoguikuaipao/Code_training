package combinationSum

func combinationSum(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	cur := make([]int, 0)
	curSum := 0
	for i, e := range candidates {
		cur = append(cur, e)
		curSum += e
		backtracking(candidates, &cur, target, &ret, curSum, i)
		curSum -= e
		cur = cur[:len(cur)-1]
	}
	return ret
}

func backtracking(candidates []int, cur *[]int, target int, ret *[][]int, curSum int, start int) {
	if curSum > target {
		return
	} else if curSum == target {
		tmp := make([]int, len(*cur))
		copy(tmp, *cur)
		*ret = append(*ret, tmp)
		return
	}

	for i, e := range candidates {
		if i < start {
			continue
		}
		*cur = append(*cur, e)
		curSum += e
		backtracking(candidates, cur, target, ret, curSum, i)
		curSum -= e
		*cur = (*cur)[:len(*cur)-1]
	}
}
