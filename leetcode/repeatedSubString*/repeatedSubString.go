package main

import "fmt"

func repeatedSubstringPattern(s string) bool {
	next := GetNext(s)
	size := len(s)
	if next[len(next)-1] != 0 && size%(size-next[len(next)-1]) == 0 {
		return true
	}
	return false
}

//	func KMP(pattern string, text string) []int {
//		next := GetNext(pattern)
//		ps := len(pattern)
//		ts := len(text)
//		idxs := make([]int, 0)
//
//		for i, j := 0, 0; i < ps && j < ts; {
//			if pattern[i] != text[j] {
//				if i > 0 {
//					i = next[i-1]
//				} else {
//					j++
//				}
//				continue
//			}
//			i++
//			j++
//			if i == ps {
//				idxs = append(idxs, j-i)
//				i = next[i-1]
//			}
//
//		}
//		return idxs
//	}

func GetNext(pattern string) (next []int) {
	size := len(pattern)
	next = make([]int, size)
	next[0] = 0
	for i, j := 1, 0; i < size; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = next[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		next[i] = j
	}
	return next
}

func main() {
	s := "ababc"
	fmt.Println(repeatedSubstringPattern(s))
}
