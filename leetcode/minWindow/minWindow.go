package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow("abcabdebac", "cda"))
}

func minWindow(s string, t string) string {
	table := make(map[byte]state)
	for _, word := range t {
		if e, ok := table[byte(word)]; ok {
			e.count++
			table[byte(word)] = e
		} else {
			table[byte(word)] = state{
				count: 1,
				now:   0,
			}
		}
	}
	list := make([]string, 0)
	start := 0
	for size := len(s); start < size; start++ {
		if _, ok := table[(s[start])]; ok {
			break
		}
	}
	min := math.MaxInt
	wStart := start
	flag := false
	for i, size := start, len(s); i < size; i++ {
		if _, ok := table[s[i]]; !ok {
			continue
		}

		current := table[s[i]]
		current.now++
		table[s[i]] = current
		if flag == false {
			iflag := false
			for _, v := range table {
				if v.now < v.count {
					iflag = true
					break
				}
			}
			if iflag == false {
				for ; wStart < i; wStart++ {
					if e, ok := table[s[wStart]]; !ok {
						continue
					} else {
						if e.now > e.count {
							e.now--
							table[s[wStart]] = e
							continue
						} else {
							break
						}
					}
				}
				if i-wStart < min {
					min = i - wStart
					list = append(list, s[wStart:i+1])
				}
				flag = true
			}
		} else {
			if s[i] == s[wStart] && current.now > current.count {
				wStart = wStart + 1
				current.now--
				table[s[i]] = current
				for ; wStart < i; wStart++ {
					if e, ok := table[s[wStart]]; !ok {
						continue
					} else {
						if e.now > e.count {
							e.now--
							table[s[wStart]] = e
							continue
						} else {
							break
						}
					}
				}
				if i-wStart < min {
					min = i - wStart
					list = append(list, s[wStart:i+1])
				}
			}
		}

	}
	if len(list) > 0 {
		return list[len(list)-1]
	}
	return ""
}

type state struct {
	count int
	now   int
}
