package main

import "fmt"

func main() {
	fmt.Println(rangeBitwiseAnd(60000, 2147483645))
}

func rangeBitwiseAnd(left int, right int) int {
	if left == 0 {
		return 0
	}
	moved := 0
	for left != right {
		left >>= 1
		right >>= 1
		moved++
	}
	return left << moved
}
