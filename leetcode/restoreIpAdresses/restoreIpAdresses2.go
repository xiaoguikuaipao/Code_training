package main

import "strconv"

func restoreIpAddresses2(s string) []string {
	ret := make([]string, 0)
	path := make([]byte, 0)

	backtracking(s, &ret, &path, 0, 0)

	return ret
}

func backtracking(s string, ret *[]string, path *[]byte, start int, dotN int) {
	if len(s) == 4 {
		tmp := make([]byte, 0)
		for i, e := range s {
			tmp = append(tmp, byte(e))
			if i != len(s)-1 {
				tmp = append(tmp, '.')
			}
		}
		*ret = append(*ret, string(tmp))
		return
	}

	if dotN > 4 {
		return
	} else if dotN == 4 && start == len(s) {
		tmp := make([]byte, len(*path)-1)
		copy(tmp, *path)
		*ret = append(*ret, string(tmp))
		return
	}
	for i := start + 1; i <= len(s); i++ {
		if isValid(s[start:i]) {
			dotN++
			*path = append(*path, []byte(s[start:i])...)
			*path = append(*path, '.')
			backtracking(s, ret, path, i, dotN)
			*path = (*path)[:len(*path)-1]
			*path = (*path)[:len(*path)-len(s[start:i])]
			dotN--
		} else {
			return
		}
	}
}

func isValid(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	ipN, err := strconv.Atoi(s)
	if err != nil || ipN > 255 {
		return false
	}
	return true
}
