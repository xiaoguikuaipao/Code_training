package main

// 节点交换规则
type selectFirst func(index, parent int) bool

type Heap[T any] struct {
	data []T
	selectFirst selectFirst
}

// 构造函数（需传入比较器）
func NewHeap[T any](selectFirst selectFirst, elements ...T) *Heap[T] {
	h := &Heap[T]{
		data: make([]T, 0, len(elements)),
		selectFirst: selectFirst,
	}
	for _, e := range elements {
		h.Insert(e)
	}
	return h
}

// 插入元素
func (h *Heap[T]) Insert(e T) {
	h.data = append(h.data, e)
	h.siftUp(len(h.data) - 1)
}

// 上浮
func (h *Heap[T]) siftUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.selectFirst(index, parent) {
			h.swap(index, parent)
			index = parent
		} else {
			break
		}
	}
}

func (h *Heap[T]) siftDown(index int) {
	size := len(h.data)
	left := 2 * index
	right := 2*index + 1
	target := left
	if h.selectFirst(index, left) {
		target = 
	}
}

// 弹出
func (h *Heap[T]) Pop() (T, bool) {
	if len(h.data) == 0 {
		var zero T
		return zero, false
	}
	ret := h.data[0]
	h.swap(0, len(h.data)-1)
	h.data = h.data[:len(h.data)-1]
	h.siftDown(0)
	return ret, true
}

// 交换节点
func (h *Heap[T]) swap(index, parent int) {
	h.data[index], h.data[parent] = h.data[parent], h.data[index]
}
