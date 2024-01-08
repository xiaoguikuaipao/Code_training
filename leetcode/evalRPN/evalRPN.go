package main

import (
	"fmt"
	"strconv"
)

func evalRPN(token []string) int {
	st := make([]string, 0)
	for _, e := range token {
		switch e {
		case "+":
			e1 := st[len(st)-1]
			e2 := st[len(st)-2]
			st = st[:len(st)-2]
			e1i, _ := strconv.Atoi(e1)
			e2i, _ := strconv.Atoi(e2)
			st = append(st, strconv.Itoa(e1i+e2i))
		case "-":
			e1 := st[len(st)-2]
			e2 := st[len(st)-1]
			st = st[:len(st)-2]
			e1i, _ := strconv.Atoi(e1)
			e2i, _ := strconv.Atoi(e2)
			st = append(st, strconv.Itoa(e1i-e2i))
		case "*":
			e1 := st[len(st)-2]
			e2 := st[len(st)-1]
			st = st[:len(st)-2]
			e1i, _ := strconv.Atoi(e1)
			e2i, _ := strconv.Atoi(e2)
			st = append(st, strconv.Itoa(e1i*e2i))
		case "/":
			e1 := st[len(st)-2]
			e2 := st[len(st)-1]
			st = st[:len(st)-2]
			e1i, _ := strconv.Atoi(e1)
			e2i, _ := strconv.Atoi(e2)
			st = append(st, strconv.Itoa(e1i/e2i))
		default:
			st = append(st, e)
		}
	}
	result, _ := strconv.Atoi(st[len(st)-1])
	return result
}

func main() {
	token := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(token))
}
