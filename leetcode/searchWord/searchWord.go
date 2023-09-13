package main

import "fmt"

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "SEE"
	fmt.Println(exist(board, word))

}

var four = [][]int{
	[]int{-1, 0},
	[]int{0, 1},
	[]int{0, -1},
	[]int{1, 0},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i, v := range board {
		for j := range v {
			if search(board, visited, 0, i, j, word) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, visited [][]bool, index, x, y int, word string) bool {
	if index == len(word)-1 {
		return board[x][y] == word[index]
	}
	if board[x][y] == word[index] {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nx := four[i][0] + x
			ny := four[i][1] + y
			if isInboard(nx, ny, board) && !visited[nx][ny] && search(board, visited, index+1, nx, ny, word) {
				return true
			}
		}
		visited[x][y] = false
	}
	return false
}

func isInboard(x, y int, board [][]byte) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}
