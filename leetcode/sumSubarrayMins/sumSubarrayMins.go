package main

import (
	"fmt"
	"io"
)

type abc interface {
	io.Writer
	sync() error
}

var ABC abc

func main() {
	arr := []int{3, 1, 2, 4, 3}
	fmt.Println(sumSubarrayMins(arr))
}

//单调栈三元素操作
func sumSubarrayMins(arr []int) int {
	n := len(arr)
	arr = append(arr, 0)
	res := 0
	mod := 1000000007
	stack := make([]int, 0, n)
	for i := 0; i <= n; i++ {
		for len(stack) > 0 && arr[i] < arr[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			res += (arr[mid] * (mid - left) * (i - mid)) % mod
		}
		stack = append(stack, i)
	}
	return res
}
