package main

import "fmt"

func KMP(pattern string, text string) {
	next := GetNext(pattern)
	ps := len(pattern)
	ts := len(text)
	for i, j := 0, 0; i < ps && j < ts; {
		if pattern[i] != text[j] {
			if i > 0 {
				// jump to next[i-1]
				i = next[i-1]
			} else {
				// i == 0 stands for the text word match nothing, so next text word.
				j++
			}
			continue
		}
		i++
		j++
		if i == ps {
			fmt.Printf("match at index %d\n", j-i)
			i = next[i-1]
		}
	}
}

// GetNext get the next array. et. "aabaaf" - next[0,1,0,1,2,0]
func GetNext(pattern string) (next []int) {
	//1. initialize the next array
	//2. prefix == suffix
	//3. prefix != suffix
	//4. update the next array
	size := len(pattern)
	next = make([]int, size)
	next[0] = 0
	//The i represents the last of the suffix, the j represents the last of the prefix.
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
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"
	KMP(pattern, text)
}
