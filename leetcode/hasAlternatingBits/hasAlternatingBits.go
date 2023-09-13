package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hasAlternatingBits(10))
}

func hasAlternatingBits(n int) bool {
	var bit = -2
	lsb := n & -n
	if n == 0 || (math.Log2(float64(lsb)) != 0 && math.Log2(float64(lsb)) != 1) {
		return false
	}
	for {
		lsb := n & -n
		if bit != -2 && math.Log2(float64(lsb)) != float64(bit+2) {
			return false
		}
		n ^= lsb
		bit = int(math.Log2(float64(lsb)))
		if n == 0 {
			return true
		}
	}
}
