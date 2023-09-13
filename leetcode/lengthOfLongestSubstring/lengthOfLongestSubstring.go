package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbbb"))
}

func lengthOfLongestSubstring(s string) int {
	path := make([]rune, 0, len(s))
	existed := make(map[rune]struct{}, len(s))
	maxL := 0
	for _, b := range s {
		if _, ok := existed[b]; !ok {
			path = append(path, b)
			existed[b] = struct{}{}
			if len(path) > maxL {
				maxL = len(path)
			}
		} else {
			for {
				if len(path) > 0 {
					delete(existed, path[0])
					path = path[1:]
				}
				if _, ok := existed[b]; !ok {
					break
				}
			}
			path = append(path, b)
			existed[b] = struct{}{}
		}
	}
	return maxL
}
