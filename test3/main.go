package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := "hello, 世界"
	for i, c := range a {
		fmt.Printf("i:%d %d\n", i, c)
		bs := make([]byte, 32)
		utf8.EncodeRune(bs, c)
		fmt.Printf("%X\n", bs)
	}
}
