package main

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] < people[j][0] {
			return false
		} else if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		} else {
			return true
		}
	})

	for i := 0; i < len(people); i++ {
		pos := people[i][1]
		if pos != i {
			temp := people[i]
			copy(people[pos+1:i+1], people[pos:i])
			people[pos] = temp
		}
	}
	return people

}
