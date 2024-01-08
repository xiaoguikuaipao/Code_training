package main

import (
	"bytes"
)

func solveNQueens2(n int) [][]string {
	ret := make([][]string, 0)
	forbidden := make([][]int, n)
	path := make([]string, 0)
	for i := range forbidden {
		forbidden[i] = make([]int, n)
	}

	backtracking(&ret, &path, n, forbidden)

	return ret
}

func backtracking(ret *[][]string, path *[]string, n int, forbidden [][]int) {
	if len(*path) == n {
		tmp := make([]string, len(*path))
		copy(tmp, *path)
		*ret = append(*ret, tmp)
		return
	}
	for i := 0; i < n; i++ {
		move := bytes.Repeat([]byte{'.'}, n)
		if forbidden[len(*path)][i] > 0 {
			continue
		}
		move[i] = 'Q'
		for x := 0; x < n; x++ {
			forbidden[x][i]++
		}
		for x, y := len(*path), i; x < len(forbidden) && y < len(forbidden[0]); {
			forbidden[x][y]++
			x = x + 1
			y = y + 1
		}
		for x, y := len(*path), i; x < len(forbidden) && y >= 0 && y < len(forbidden[0]); {
			forbidden[x][y]++
			x = x + 1
			y = y - 1
		}
		*path = append(*path, string(move))

		backtracking(ret, path, n, forbidden)

		*path = (*path)[:len(*path)-1]
		for x := 0; x < n; x++ {
			forbidden[x][i]--
		}
		for x, y := len(*path), i; x < len(forbidden) && y < len(forbidden[0]); {
			forbidden[x][y]--
			x = x + 1
			y = y + 1
		}
		for x, y := len(*path), i; x < len(forbidden) && y >= 0 && y < len(forbidden[0]); {
			forbidden[x][y]--
			x = x + 1
			y = y - 1
		}
	}
}

//func main() {
//	fmt.Println(solveNQueens2(4))
//}
