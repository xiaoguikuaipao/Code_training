package main

import "fmt"

func main() {
	n := 4
	fmt.Println(solveNQueens(n))
}

var res [][]string

func solveNQueens(n int) [][]string {
	res = make([][]string, 0, 100)
	board := make([][]byte, n)
	restraint := make([][]int, n)
	for i := range board {
		newRow := make([]byte, n)
		for ii := range newRow {
			newRow[ii] = '.'
		}
		board[i] = newRow
	}
	for i := range board {
		newRow := make([]int, n)
		restraint[i] = newRow
	}
	dfs(board, restraint, 0)
	return res
}

func dfs(board [][]byte, restraint [][]int, row int) {
	n := len(board)
	if row == n {
		return
	}
	for col, _ := range board[row] {
		if restraint[row][col] == 0 && board[row][col] == '.' {
			board[row][col] = 'Q'
			if row == n-1 {
				store := make([][]byte, n)
				copy(store, board)
				var s []string
				for _, v := range store {
					s = append(s, string(v))
				}
				res = append(res, s)
			}
			for rcol := range restraint[row] {
				restraint[row][rcol]++
			}
			for i := 0; i < n; i++ {
				restraint[i][col]++
			}
			for lrow, lcol := row-1, col-1; lrow >= 0 && lcol >= 0; {
				restraint[lrow][lcol]++
				lrow--
				lcol--
			}
			for lrow, lcol := row-1, col+1; lrow >= 0 && lcol < n; {
				restraint[lrow][lcol]++
				lrow--
				lcol++
			}
			for lrow, lcol := row+1, col-1; lrow < n && lcol >= 0; {
				restraint[lrow][lcol]++
				lrow++
				lcol--
			}
			for lrow, lcol := row+1, col+1; lrow < n && lcol < n; {
				restraint[lrow][lcol]++
				lrow++
				lcol++
			}
			dfs(board, restraint, row+1)
			board[row][col] = '.'
			for rcol := range restraint[row] {
				restraint[row][rcol]--
			}
			for i := 0; i < n; i++ {
				restraint[i][col]--
			}
			for lrow, lcol := row-1, col-1; lrow >= 0 && lcol >= 0; {
				restraint[lrow][lcol]--
				lrow--
				lcol--
			}
			for lrow, lcol := row+1, col+1; lrow < n && lcol < n; {
				restraint[lrow][lcol]--
				lrow++
				lcol++
			}
			for lrow, lcol := row-1, col+1; lrow >= 0 && lcol < n; {
				restraint[lrow][lcol]--
				lrow--
				lcol++
			}
			for lrow, lcol := row+1, col-1; lrow < n && lcol >= 0; {
				restraint[lrow][lcol]--
				lrow++
				lcol--
			}
		}
	}
}
