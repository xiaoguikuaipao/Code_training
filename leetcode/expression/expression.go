package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compute("114.514+2023*6/11-1919/810"))
}
func compute(data string) float64 {
	// 在这⾥写代码
	now := 0
	stack := make([]string, 0)
	size := len(data)
	for i := 0; i <= size; i++ {

		if i == size || data[i] == '+' || data[i] == '-' || data[i] == '*' || data[i] == '/' {
			stack = append(stack, data[now:i])
			for len(stack) > 2 && (stack[len(stack)-2] == "*" || stack[len(stack)-2] == "/") {
				op1, _ := strconv.ParseFloat(stack[len(stack)-1], 32)
				stack = stack[:len(stack)-1]
				op2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				op3, _ := strconv.ParseFloat(stack[len(stack)-1], 32)
				stack = stack[:len(stack)-1]
				res := 0.0
				switch op2 {
				case "+":
					res = op3 + op1
				case "-":
					res = op3 - op1
				case "*":
					res = op3 * op1
				case "/":
					res = op3 / op1
				}
				stack = append(stack, strconv.FormatFloat(res, 'f', -1, 32))
			}
			if i < size {
				stack = append(stack, string(data[i]))
				now = i + 1
			}
		}
	}
	for len(stack) > 2 {
		op1, _ := strconv.ParseFloat(stack[len(stack)-1], 32)
		stack = stack[:len(stack)-1]
		op2 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		op3, _ := strconv.ParseFloat(stack[len(stack)-1], 32)
		stack = stack[:len(stack)-1]
		res := 0.0
		switch op2 {
		case "+":
			res = op3 + op1
		case "-":
			res = op3 - op1
		}
		stack = append(stack, strconv.FormatFloat(res, 'f', -1, 32))
	}
	ret, _ := strconv.ParseFloat(stack[len(stack)-1], 64)
	return ret
}
