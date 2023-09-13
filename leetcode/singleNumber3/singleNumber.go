package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}

func singleNumber(nums []int) []int {
	res := make([]int, 2)
	final := 0
	for _, v := range nums {
		final ^= v
	}
	// 小技巧：可以获得一个数的最低有效位，或者判断该数是否为2的幂
	// 本题中，最低有效位为1，说明该位的产生源头是0^=1，用此最低有效位&这两个数结果分别是0和1，将此作为if条件分流提取结果
	final &= -final
	for _, v := range nums {
		if final&v == 0 {
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}
	return res
}
