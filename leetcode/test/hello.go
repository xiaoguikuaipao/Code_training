package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compress("adef"))
}

func compress(str string) string {
	res := make([]byte, 0)
	size := len(str)
	var count int
	for i := 0; i < size; i++ {
		if i > 0 && str[i] == str[i-1] {
			count++
			continue
		}
		if count > 1 {
			res = append(res, []byte(strconv.Itoa(count))...)
			count = 0
		} else {
			count = 0
		}
		res = append(res, str[i])
		count++
	}
	if count > 1 {
		res = append(res, []byte(strconv.Itoa(count))...)
	}
	return string(res)
}
