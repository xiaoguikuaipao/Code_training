package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 9))
}

var res [][]int

func combinationSum3(k int, n int) [][]int {
	res = make([][]int, 0)
	path := make([]int, 0, k)
	for i := 1; i < 10; i++ {
		path = append(path, i)
		dfs(i, path, n, i, k)
		path = path[:len(path)-1]
	}
	return res
}

func dfs(now int, path []int, sum int, currentSum int, k int) {
	if len(path) == k {
		if currentSum == sum {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		return
	}
	for i := now + 1; i < 10; i++ {
		path = append(path, i)
		dfs(i, path, sum, currentSum+i, k)
		path = path[:len(path)-1]
	}
}
