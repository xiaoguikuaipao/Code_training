package main

import "fmt"

func main() {
	fmt.Println(MajorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

func MajorityElement(nums []int) int {
	n := len(nums)
	record := make(map[int]int, n/2)
	for _, num := range nums {
		record[num] += 1
		if record[num] > n/2 {
			return num
		}
	}
	return 0
}
