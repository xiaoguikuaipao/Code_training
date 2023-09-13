package main

import "fmt"

func main() {
	nums := []int{3, 4, 2, 1}
	nums = pancakeSort(nums)
	fmt.Println(nums)
}

var res []int

func pancakeSort(arr []int) []int {
	res = make([]int, 0, 10*len(arr))
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		if arr[i] == i+1 {
			continue
		} else {
			maxIdx := find(arr, i)
			reverse(arr, maxIdx)
			reverse(arr, i)
		}
	}

	return res
}

func reverse(arr []int, idx int) {
	for i, j := 0, idx; i < j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	res = append(res, idx+1)
}
func find(arr []int, n int) int {
	for i, v := range arr[:n] {
		if v == n+1 {
			return i
		}
	}
	return -1
}
