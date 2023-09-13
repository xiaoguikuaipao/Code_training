package main

import (
	"fmt"
)

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}

var res []int

func findSubstring(s string, words []string) []int {
	res = make([]int, 0)
	table := make(map[string]int)
	nSub := len(words)
	subLen := len(words[0])
	for _, word := range words {
		table[word] = 0
	}
	for _, word := range words {
		table[word]++
	}
	for i := 0; i+subLen*nSub <= len(s); i++ {
		if _, ok := table[s[i:i+subLen]]; ok {
			if forwardSearch(table, s, i, words) {
				res = append(res, i)
			}
			for _, word := range words {
				table[word] = 0
			}
			for _, word := range words {
				table[word]++
			}
		}
	}
	return res
}

func forwardSearch(table map[string]int, s string, start int, words []string) bool {
	subLen := len(words[0])
	nSub := len(words)
	if start+nSub*subLen > len(s) {
		return false
	}
	table[s[start:start+subLen]] -= 1
	for start+2*subLen <= len(s) {
		start += subLen
		if c, ok := table[s[start:start+subLen]]; ok && c > 0 {
			table[s[start:start+subLen]] -= 1
		} else {
			break
		}
	}
	for _, v := range table {
		if v != 0 {
			return false
		}
	}
	return true
}
