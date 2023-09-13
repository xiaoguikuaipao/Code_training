package main

import "fmt"

func main() {
	heights := []int{
		2, 1, 5, 6, 2, 3, 3, 3,
	}
	fmt.Println(largestRectangleArea(heights))
}

func largestRectangleArea(heights []int) int {
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	res := 0
	n := len(heights)
	st := make([]int, 0, n/2)
	st = append(st, 0)
	for i := 1; i < n; i++ {
		if heights[i] > heights[st[len(st)-1]] {
			st = append(st, i)
		} else {
			for len(st) > 1 && heights[i] < heights[st[len(st)-1]] {
				right := i
				mid := heights[st[len(st)-1]]
				st = st[:len(st)-1]
				left := st[len(st)-1]
				res = max((right-left-1)*mid, res)
			}
			st = append(st, i)
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
