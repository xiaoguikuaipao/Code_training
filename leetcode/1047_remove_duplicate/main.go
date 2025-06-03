package main

import (
	"fmt"
)

/*
示例：
输入："abbaca"
输出："ca"
解释：
例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。
*/

type Stacker[T any] interface {
	Pop() T
	Push(T)
	Empty() bool
	Peak() T
	Show() []T
	Bottom() T
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Pop() T {
	if s.Empty() {
		panic("empty stack")
	}
	result := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return result
}

func (s *Stack[T]) Push(item T) {
	if s.items == nil {
		s.items = make([]T, 0)
	}
	s.items = append(s.items, item)
}

func (s *Stack[T]) Peak() T {
	if s.Empty() {
		panic("empty stack")
	}
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Show() []T {
	return s.items
}

func (s *Stack[T]) Bottom() T {
	return s.items[0]
}

func NewStacker[T any](cap ...int) Stacker[T] {
	st := &Stack[T]{
		items: make([]T, 0),
	}
	if len(cap) == 1 {
		st.items = make([]T, 0, cap[0])
	}
	return st
}

func removeDuplicates(s string) string {
	st := NewStacker[rune](len(s))
	for _, c := range s {
		if st.Empty() {
			st.Push(c)
			continue
		}
		e := st.Peak()
		if e == c {
			_ = st.Pop()
		} else {
			st.Push(c)
		}
	}
	return string(st.Show())
}

func main() {
	s := "abbaca"
	fmt.Println(removeDuplicates(s))
}
