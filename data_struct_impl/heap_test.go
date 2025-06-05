package datastructimpl

import (
	"testing"
)

type ts struct {
	num   int
	other string
}

func TestHeap(t *testing.T) {
	tss := make([]ts, 0)
	for i := range 15 {
		tss = append(tss, ts{num: i})
	}
	h := NewHeap(func(index, parent int) bool { return tss[index].num > tss[parent].num }, tss...)
	t.Logf("%v", h.Data())
}
