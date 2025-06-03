package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	result := make(map[int]int)
	result[0] = 0
	result[1] = 0
	for i := 0; i < 1000; i++ {
		store := make([]int, 10000)
		record := make([]int, 0)
		max := big.NewInt(100)
		count := 0
		for j := 0; j < 10000; j++ {
			if count == 15 {
				record = append(record, j)
				count = 0
			}
			c, err := rand.Int(rand.Reader, max)
			if err != nil {
				return
			}
			if c.Int64() < 50 {
				store[j] = 0
				count++
			} else {
				store[j] = 1
				count = 0
			}
		}
		for _, k := range record {
			result[store[k]]++
		}
		fmt.Printf("The %d exp[12次连续0后的第13位所在索引]:%+v\n", i, record)
	}
	fmt.Printf("%+v\n", result)
}
