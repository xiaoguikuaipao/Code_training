package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(reverseBits(0b00000010100101000001111010011100))
}

func reverseBits(num uint32) uint32 {
	return bits.Reverse32(num)
}
