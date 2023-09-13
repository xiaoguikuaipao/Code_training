package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var Map = make([][]string, 30)
var Seen = make([][]int, 30)
var d = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
var flag = 0
var ans = 0

func main() {
	for i := range Map {
		Map[i] = make([]string, 30)
	}
	for i := range Seen {
		Seen[i] = make([]int, 30)
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入行数：")
	scanner.Scan()
	_, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("无效输入:%v", err)
	}
	fmt.Println("请输入列数：")
	scanner.Scan()
	_, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("无效输入:%v", err)
	}
	fmt.Println("请输入矩阵：（格式为11,12;21,22）")
	scanner.Scan()
	line := scanner.Text()
	row := strings.Split(line, ";")
	var ti, tj int
	for i, everyRow := range row {
		singleRow := strings.Split(everyRow, ",")
		for j, every := range singleRow {
			Map[i][j] = every
			fmt.Print(Map[i][j])
			if j == len(singleRow)-1 {
				fmt.Println()
			}
			ti = i
			tj = j
		}
	}
	for i := 1; i < ti; i++ {
		for j := 1; j < tj; j++ {
			if Map[i][j] == "#" && Seen[i][j] == 0 {
				flag = 0
				dfs(i, j)
				if flag == 0 {
					ans++
				}
			}
		}
	}
	fmt.Printf("有%d座岛被淹没", ans)
}

func dfs(x, y int) {
	Seen[x][y] = 1
	if Map[x][y+1] == "#" && Map[x][y-1] == "#" && Map[x+1][y] == "#" && Map[x-1][y] == "#" {
		flag = 1
	}
	for i := 0; i < 4; i++ {
		nx := x + d[i][0]
		ny := y + d[i][1]
		if Seen[nx][ny] == 0 && Map[nx][ny] == "#" {
			dfs(nx, ny)
		}
	}
}
