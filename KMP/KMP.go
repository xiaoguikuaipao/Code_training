package main

import (
	"fmt"
)

func GetTabPrefix(pattern string) []int {
	prefix := make([]int, len(pattern))
	for lp, ls := 0, 1; ls < len(pattern); ls++ {
		if pattern[lp] == pattern[ls] {
			lp++
		} else {
			for lp > 0 && pattern[lp] != pattern[ls] {
				lp = prefix[lp-1]
			}
			if pattern[lp] == pattern[ls] {
				lp++
			}
		}
		prefix[ls] = lp
	}
	return prefix
}

func KMP(s, pattern string) []int {
	prefix := GetTabPrefix(pattern)
	idxs := make([]int, 0)
	ip := 0
	for is, cs := range []byte(s) {
		if pattern[ip] != cs {
			for pattern[ip] != cs && ip > 0 {
				ip = prefix[ip-1]
			}
		}
		if pattern[ip] == cs {
			ip++
			if ip == len(pattern) {
				idxs = append(idxs, is-len(pattern)+1)
				ip = 0
			}
		}
	}
	return idxs
}

func main() {
	// text := "ABABDABACDABABCABAB"
	// pattern := "ABABCABAB"
	text := "你好我好大家好你好我好大家坏"
	pattern := "你好"
	idxs := KMP(text, pattern)
	for _, id := range idxs {
		fmt.Println(id, ":", text[id:(id+len(pattern))])
	}
	fmt.Println(idxs)
}
