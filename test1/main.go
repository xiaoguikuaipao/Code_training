package main

import (
	"fmt"
	"reflect"
)

func main() {
	test := []byte{'1', '2', '3'}
	fmt.Println(test)
	var y myint = []int{1, 2, 3}
	ptr := &y
	ry := reflect.ValueOf(y).Type()
	py := reflect.ValueOf(&y).Elem()
	fmt.Println(ry.String(), py.String())
	methodValue, _ := ry.MethodByName("print")
	fmt.Println(methodValue)
	fmt.Printf("%p", ptr)
	var empty interface{}
	empty = y
	fmt.Printf("%p", empty)
	fmt.Printf("%d\n", (-10+9)%9)
	fmt.Println(10 ^ 9)
}

type myint []int

func (mi myint) print() {
	fmt.Println(mi[0])
}
