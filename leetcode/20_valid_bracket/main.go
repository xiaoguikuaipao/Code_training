package main

import (
	"errors"
	"fmt"
)

/*
示例 1：
输入：s = "()"
输出：true

示例 2：
输入：s = "()[]{}"
输出：true
*/

var (
	ErrStackEmpty = errors.New("stack is empty")
)

type Stacker[T any] interface {
	Pop() (T, error)
	Push(T)
	Empty() bool
	Peak() (T, error)
}

type Stack[T any] struct {
	Items []T
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Empty() {
		var zero T
		return zero, ErrStackEmpty
	}
	result := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return result, nil
}

func (s *Stack[T]) Push(item T) {
	if s.Items == nil {
		s.Items = make([]T, 0)
	}
	s.Items = append(s.Items, item)
}

func (s *Stack[T]) Peak() (T, error) {
	if s.Empty() {
		var zero T
		return zero, ErrStackEmpty
	}
	return s.Items[len(s.Items)-1], nil
}

func (s *Stack[T]) Empty() bool {
	return len(s.Items) == 0
}

func NewStacker[T any](cap ...int) Stacker[T] {
	st := &Stack[T]{
		Items: make([]T, 0),
	}
	if len(cap) == 1 {
		st.Items = make([]T, 0, cap[0])
	}
	return st
}

func isValid(s string) bool {
	st := NewStacker[rune](len(s))
	for _, c := range s {
		switch string(c) {
		case "{", "(", "[":
			st.Push(c)
		case "}", ")", "]":
			e, err := st.Peak()
			if err != nil {
				return false
			}
			if c == '}' && e == '{' {
				_, _ = st.Pop()
			} else if c == ']' && e == '[' {
				_, _ = st.Pop()
			} else if c == ')' && e == '(' {
				_, _ = st.Pop()
			} else {
				return false
			}
		}
	}
	return st.Empty()
}

func main() {
	s := "(])"
	fmt.Println(isValid(s))
}
