package main

import "fmt"

func main() {
	str := "a"
	fmt.Println(partition(str))
}

var res [][]string

func partition(s string) [][]string {
	res = make([][]string, 0, 100)
	n := len(s)
	for i := 0; i < n; i++ {
		part := make([]int, 0, n)
		dfs(s, i+1, part)
	}
	return res
}

func dfs(s string, sep int, part []int) {
	n := len(s)
	last := 0
	if len(part) > 0 {
		last = part[len(part)-1]
	}
	sub := s[last:sep]
	if check(sub) {
		part = append(part, sep)
		if len(part) > 0 && part[len(part)-1] == n {
			tmp := make([]string, 0, n)
			for i := 0; i < len(part); i++ {
				if i == 0 {
					tmp = append(tmp, s[0:part[i]])
				} else {
					tmp = append(tmp, s[part[i-1]:part[i]])
				}
			}
			res = append(res, tmp)
		}
		for nextsep := sep + 1; nextsep <= n; nextsep++ {
			dfs(s, nextsep, part)
		}
		part = part[:len(part)-1]
	}
}

func check(s string) bool {
	end := len(s) - 1
	begin := 0
	for begin < end {
		if s[begin] == s[end] {
			begin++
			end--
		} else {
			break
		}
	}
	if begin == end || begin-1 == end {
		return true
	}
	return false
}
