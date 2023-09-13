package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var stickLen = make([]int, 0, 20)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入木棍数:")
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("输入有误: %v", err)
	}
	fmt.Println("请输入每根长度: ")
	for len(stickLen) < n {
		scanner.Scan()
		Read := strings.Split(scanner.Text(), " ")
		for _, e := range Read {
			length, err := strconv.Atoi(e)
			if err != nil {
				log.Fatalf("输入错误: %v", err)
			}
			stickLen = append(stickLen, length)
		}
	}
	solve(n)
}

func solve(n int) {
	Total := 0
	for i := 0; i < n; i++ {
		Total += stickLen[i]
	}
	divisor := divisor(Total)
	for i := range divisor {
		dfs(i)
	}
}

func dfs(singleLen int) {

}

func divisor(x int) []int {
	divisor := make([]int, 0)
	for i := 1; i < int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			divisor = append(divisor, i)
		}
	}
	return divisor
}
