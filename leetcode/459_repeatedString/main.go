package main

import "fmt"

// 示例 1:
// 输入: s = "abab"
// 输出: true
// 解释: 可由子串 "ab" 重复两次构成。

func repeatedSubstringPattern(s string) bool {
	prefix := GetTabPrefix(s)
	l := len(prefix)
	if prefix[l-1] != 0 && len(s)%(len(s)-prefix[l-1]) == 0 {
		return true
	} else {
		return false
	}
}

func GetTabPrefix(pattern string) []int {
	prefix := make([]int, len(pattern))
	for lp, ls := 0, 1; ls < len(prefix); ls++ {
		if pattern[lp] == pattern[ls] {
			lp++
		} else {
			for pattern[lp] != pattern[ls] && lp > 0 {
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

func main() {
	s := "abacababacab"
	fmt.Println(repeatedSubstringPattern(s))
}
