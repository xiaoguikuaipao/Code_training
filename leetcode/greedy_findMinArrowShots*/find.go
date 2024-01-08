package greedy_findMinArrowShots

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	sum := 0

	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] {
			sum++
		} else {
			points[i][1] = min(points[i-1][1], points[i][1])
		}
	}
	return sum

}
