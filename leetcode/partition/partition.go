package main

import "fmt"

func partition(s string) [][]string {
	speedUp := make(map[string]struct{})
	ret := make([][]string, 0)
	path := make([]string, 0)
	backtracking(s, &ret, &path, 0, speedUp)
	return ret
}

func backtracking(s string, ret *[][]string, path *[]string, start int, speedUp map[string]struct{}) {
	if start == len(s) {
		tmp := make([]string, len(*path))
		copy(tmp, *path)
		*ret = append(*ret, tmp)
		return
	}
	for i := start + 1; i <= len(s); i++ {
		if isPalindrome(s[start:i], speedUp) {
			*path = append(*path, s[start:i])
			backtracking(s, ret, path, i, speedUp)
			*path = (*path)[:len(*path)-1]
		}
	}
}

func isPalindrome(s string, speedUp map[string]struct{}) bool {
	if _, ok := speedUp[s]; ok {
		return true
	}
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	speedUp[s] = struct{}{}
	return true
}

func main() {
	fmt.Println(partition("aab"))
}
