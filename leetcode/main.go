package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var b []byte
	hello := make(map[string]interface{})
	hello["1"] = 2
	s := "哈哈，你是xiaokeai"
	b = []byte(s)
	var r []rune
	r = []rune(s)
	_, s2 := utf8.DecodeRune(b)
	_, s3 := utf8.DecodeRuneInString(s)
	fmt.Println(len(b), len(r), s2, s3) // 23, 13, 3, 3 , 23 = 5*3(5个中文字符) + 8*1(8个英文ASCII) ，3是一个中文rune所占字节数
	a := []int{6, 9, 2, 3, 4, 7, 5, 1, 8}
	str := fmt.Sprint(a)
	fmt.Println(str)
}
