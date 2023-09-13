package main

import (
	"fmt"
	"math"
)

func main() {
	n := 0
	k := 0
	fmt.Scan(&n, &k)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println(maxAfterDel(n, k, array))
}

func maxAfterDel(n, k int, array []int) int {
	fin := math.MinInt
	for i := 0; i < n-k; i++ {
		m := math.MaxInt
		s := 0
		for j := i; j < i+k+1; j++ {
			if array[j] < m {
				m = array[j]
			}
			s += array[j]
		}
		if s-m > fin {
			fin = s - m
		}
	}
	return fin
}
