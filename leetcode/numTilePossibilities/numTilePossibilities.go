package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(numTilePossibilities("AAB"))
}

var res []string

func numTilePossibilities(tiles string) int {
	res = make([]string, 0, 1000)
	used := make([]int, len(tiles))
	path := make([]string, 0, len(tiles))
	sortTiles := []rune(tiles)
	sort.Slice(sortTiles, func(i, j int) bool {
		return sortTiles[i] < sortTiles[j]
	})
	tiles = string(sortTiles)
	dfs(tiles, used, path)
	return len(res)
}

func dfs(tiles string, used []int, path []string) {
	for i, v := range tiles {
		if used[i] == 1 {
			continue
		}
		if i > 0 && tiles[i] == tiles[i-1] && used[i-1] == 0 {
			continue
		}
		used[i] = 1
		path = append(path, string(v))
		tmp := strings.Join(path, "")
		res = append(res, tmp)
		dfs(tiles, used, path)
		used[i] = 0
		path = path[:len(path)-1]
	}
}
