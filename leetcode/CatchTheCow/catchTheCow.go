package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var xOfCow int

type state struct {
	xf    int
	xc    int
	depth int
}

var queue = make([]*state, 0, 20)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入农夫位置:")
	scanner.Scan()
	xOfFarmer, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("Error input: %v", err)
	}
	fmt.Println("请输入牛的位置:")
	scanner.Scan()
	xOfCow, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("Error input: %v", err)
	}
	queue = append(queue, &state{
		xf:    xOfFarmer,
		xc:    xOfCow,
		depth: 0,
	})
	solve()
}

func solve() {
	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]
		if now.xf == now.xc {
			fmt.Printf("已到达,花费步数:%d", now.depth)
			return
		} else if now.xf > now.xc {
			queue = append(queue, &state{
				xf:    now.xf - 1,
				xc:    now.xc,
				depth: now.depth + 1,
			})
		} else if now.xf < now.xc {
			queue = append(queue, &state{
				xf:    now.xf + 1,
				xc:    now.xc,
				depth: now.depth + 1,
			})
			queue = append(queue, &state{
				xf:    now.xf * 2,
				xc:    now.xc,
				depth: now.depth + 1,
			})
			queue = append(queue, &state{
				xf:    now.xf - 1,
				xc:    now.xc,
				depth: now.depth + 1,
			})
		}
	}
}
