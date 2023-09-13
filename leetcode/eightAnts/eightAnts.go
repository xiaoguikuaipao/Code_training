package main

import "fmt"

var queue = make([]node, 0, 100)
var Seen = make(map[string]struct{})

type node struct {
	s    string
	step int
}

func main() {
	s := "012345678"
	queue = append(queue, node{
		s:    s,
		step: 0,
	})
	Seen[s] = struct{}{}
	solve()
}

func solve() {
	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]
		s := now.s
		step := now.step
		if s == "087654321" {
			fmt.Printf("跳了%d步", step)
			return
		}
		var i int
		for i = 0; i < len(s); i++ {
			if s[i] == '0' {
				break
			}
		}
		for j := i - 2; j <= i+2; j++ {
			k := (j + 9) % 9
			if k == i {
				continue
			}
			runes := []rune(s)
			runes[i], runes[k] = runes[k], runes[i]
			s = string(runes)
			if _, ok := Seen[s]; !ok {
				Seen[s] = struct{}{}
				queue = append(queue, node{
					s:    s,
					step: step + 1,
				})
			}
		}
	}
}
