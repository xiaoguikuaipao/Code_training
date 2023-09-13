package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	fmt.Println(board)
}

func solveSudoku(board [][]byte) {
	for row, v := range board {
		for col, literal := range v {
			if literal == '.' {
				PossibleValue := getPossibleValue(board, row, col)
				for _, value := range PossibleValue {
					board[row][col] = byte(value) + '0'
					if !dfs(board) {
						board[row][col] = '.'
					} else {
						return
					}
				}
			}
		}
	}
}
func dfs(board [][]byte) bool {
	NotFinished := false
	for row, v := range board {
		for col, literal := range v {
			if literal == '.' {
				NotFinished = true
				PossibleValue := getPossibleValue(board, row, col)
				if len(PossibleValue) == 0 {
					return false
				}
				for _, value := range PossibleValue {
					board[row][col] = byte(value) + '0'
					if !dfs(board) {
						board[row][col] = '.'
					} else {
						return true
					}
				}
			}
		}
	}
	return !NotFinished
}

func getPossibleValue(board [][]byte, row, col int) []int {
	value := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	existed := make([]int, 0, 18)
	for _, literal := range board[row] {
		if literal != '.' {
			existed = append(existed, int(literal-'0'))
		}
	}
	for i := 0; i < 9; i++ {
		if board[i][col] != '.' {
			existed = append(existed, int(board[i][col]-'0'))
		}
	}
	blockPosition := getBlockPosition(row, col)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			v := board[(blockPosition/3)*3+i][blockPosition%3*3+j]
			if v != '.' {
				existed = append(existed, int(v-'0'))
			}
		}
	}
	for _, v := range existed {
		for i, p := range value {
			if p == v {
				value = append(value[:i], value[i+1:]...)
			}
		}
	}
	return value
}

func getBlockPosition(row, col int) int {

	return row/3*3 + col/3
}
