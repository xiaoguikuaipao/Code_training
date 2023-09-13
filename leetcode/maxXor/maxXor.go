package main

import "fmt"

func main() {
	fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
}

func findMaximumXOR(nums []int) int {
	currentResult := 0
	mask := 0
	for i := 31; i >= 0; i-- {
		mask |= 1 << i
		table := make(map[int]bool, len(nums))
		for _, num := range nums {
			table[mask&num] = true
		}
		forwardResult := currentResult | (1 << i)
		for key := range table {
			// key是A，forwardResult是B，key^forwardResult是C，如果用假设的B去^A得到table里的C，说明B可以通过A^C得到
			if table[key^forwardResult] == true {
				currentResult = forwardResult
				break
			}
		}
	}
	return currentResult
}
