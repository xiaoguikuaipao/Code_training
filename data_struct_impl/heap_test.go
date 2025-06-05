package datastructimpl

import (
	"fmt"
	"testing"
)

type ts struct {
	num   int
	other string
}

func TestHeap(t *testing.T) {
	tss := make([]ts, 0)
	for i := 0; i < 15; i++ {
		tss = append(tss, ts{num: i})
	}
	h := NewHeap(func(index, parent int) bool { return tss[index].num > tss[parent].num }, tss)
	fmt.Printf("%v", h.Data())
}
