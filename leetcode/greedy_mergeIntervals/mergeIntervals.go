package greedy_mergeIntervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ret := make([][]int, 0)
	start := intervals[0][0]
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
		} else {
			ret = append(ret, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}
	ret = append(ret, []int{start, end})
	return ret
}
