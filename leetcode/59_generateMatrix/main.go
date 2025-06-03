package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(4))
}

//给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
// 示例 1：
// 输入：n = 3
// 输出：[[1,2,3],[8,9,4],[7,6,5]]

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for row := range matrix {
		matrix[row] = make([]int, n)
	}
	count := 1
	cycle := n / 2
	for c := 0; c < cycle; c++ {
		startX, startY := c, c
		x, y := startX, startY
		for ; x < n-c-1; x++ {
			matrix[y][x] = count
			count++
		}
		for ; y < n-c-1; y++ {
			matrix[y][x] = count
			count++
		}
		for ; x > startX; x-- {
			matrix[y][x] = count
			count++
		}
		for ; y > startY; y-- {
			matrix[y][x] = count
			count++
		}
	}
	if n%2 == 1 {
		matrix[n/2][n/2] = count
	}
	return matrix
}
