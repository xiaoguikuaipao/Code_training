package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(singleNumber([]int{0, 0, 0, 1, 1, 1, 88}))
}

func singleNumber(nums []int) int {
	sum := make([]byte, 64)
	for _, num := range nums {
		flag := 0
		if num > 0 {
			flag = 1
		} else {
			num = -num
		}
		b := strconv.FormatInt(int64(num), 2)
		l := len(b)
		for i, v := range b {
			if v == '1' {
				sum[l-i-1] += 1
				sum[l-i-1] %= 3
			}
		}
		if flag == 0 {
			sum[63] += 1
			sum[63] %= 3
		}
	}
	var res float64
	res = 0
	for i := 0; i < 63; i++ {
		if sum[i] == 1 {
			res += math.Pow(2, float64(i))
		}
	}
	if sum[63] == 1 {
		res = -res
	}
	return int(res)
}
