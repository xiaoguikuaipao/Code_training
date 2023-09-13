package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isAdditiveNumber("12122436"))
}

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	queue := make([]byte, 0, len(num))
	i := 0
	for i < 3 {
		queue = append(queue, num[i])
		i++
	}
	if slide(&queue, 0, 1, 2, 2, num, false) {
		return true
	}
	return false
}

func slide(queue *[]byte, n1, n2, nR int, i int, num string, next bool) bool {
	if nR == len(num) {
		return true
	}
	for i <= len(num) {
		NeedExpand := true
		n := len(*queue)
		for nr := nR; nr < n && n-nr >= min(nr-n2, n2-n1); {
			for b2 := n2; b2 < nr; {
				b1 := n1
				if p, ok := check(queue, b1, b2, nr); ok {
					NeedExpand = false
					newQ := make([]byte, len(*queue))
					copy(newQ, *queue)
					if slide(&newQ, p[1], p[2], i+1, i, num, true) {
						return true
					} else {
						NeedExpand = true
					}
				}
				if next == false {
					b2++
				} else {
					break
				}
			}
			if next == false {
				nr++
			} else {
				break
			}
		}
		if NeedExpand {
			i++
			if i < len(num) {
				*queue = append(*queue, num[i])
			}
		}
	}
	return false
}

func check(queue *[]byte, n1, n2, nR int) ([]int, bool) {
	n := len(*queue)
	c1 := n2 - n1
	c2 := nR - n2
	cR := n - nR
	N1, N2, NR := 0, 0, 0
	if c1 > 1 && int((*queue)[n1]-'0') == 0 {
		return []int{}, false
	}
	if c2 > 1 && int((*queue)[n2]-'0') == 0 {
		return []int{}, false
	}
	if cR > 1 && int((*queue)[nR]-'0') == 0 {
		return []int{}, false
	}
	for i := 0; i < c1; i++ {
		N1 += int((*queue)[n1+i]-'0') * int(math.Pow10(c1-i-1))
	}
	for i := 0; i < c2; i++ {
		N2 += int((*queue)[n2+i]-'0') * int(math.Pow10(c2-i-1))
	}
	for i := 0; i < cR; i++ {
		NR += int((*queue)[nR+i]-'0') * int(math.Pow10(cR-i-1))
	}
	if N1+N2 == NR {
		return []int{n1, n2, nR}, true
	} else {
		return []int{}, false
	}
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
