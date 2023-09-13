package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(restoreIpAddresses("010010"))
}

var res []string

func restoreIpAddresses(s string) []string {
	res = make([]string, 0, 20)
	n := len(s)
	path := make([]byte, 0, n+3)
	path = append(path, s[0])
	dfs(s, 1, path)
	return res
}

func dfs(s string, index int, path []byte) {
	if check(path) {
		if index == len(s) {
			count := 0
			for _, v := range path {
				if v == '.' {
					count++
				}
			}
			if count == 3 {
				res = append(res, string(path))
			}
		} else {
			path = append(path, s[index])
			dfs(s, index+1, path)
			path = path[:len(path)-1]

			path = append(path, '.')
			path = append(path, s[index])
			dfs(s, index+1, path)
			path = path[:len(path)-2]
		}
	} else {
		return
	}
}

func check(path []byte) bool {
	n := len(path)
	start := n - 1
	count := 0
	if start == 0 {
		return true
	} else {
		sum := 0
		for start >= 0 && path[start] != '.' {
			sum += int(path[start]-'0') * int(math.Pow10(count))
			start--
			count++
		}
		if sum > 255 {
			return false
		}
		if start > 0 && path[start] == '.' && path[start+1] == '0' && start+2 < n {
			return false
		}
		if start <= 0 && path[0] == '0' && n > 1 {
			return false
		}
	}
	return true
}
