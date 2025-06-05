package datastructimpl

// 节点交换规则
type IsSiftUp func(index, parent int) bool

type Heap[T any] struct {
	data     []T
	isSiftUp IsSiftUp
}

// 构造函数（需传入比较器）
func NewHeap[T any](isSiftUp IsSiftUp, elements ...T) *Heap[T] {
	h := &Heap[T]{
		data:     make([]T, 0, len(elements)),
		isSiftUp: isSiftUp,
	}
	for _, e := range elements {
		h.Insert(e)
	}
	h.heapify()
	return h
}

func (h *Heap[T]) Data() []T {
	return h.data
}

// O(n)
func (h *Heap[T]) heapify() {
	lastNonLeaf := len(h.data)/2 - 1
	for i := lastNonLeaf; i >= 0; i-- {
		h.siftDown(i)
	}
}

// 插入元素 O(1)
func (h *Heap[T]) Insert(e T) {
	h.data = append(h.data, e)
	h.siftUp(len(h.data) - 1)
}

// 上浮 O(1)
func (h *Heap[T]) siftUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.isSiftUp(index, parent) {
			h.swap(index, parent)
			index = parent
		} else {
			break
		}
	}
}

func (h *Heap[T]) siftDown(index int) {
	lastNonLeaf := len(h.data)/2 - 1
	for index <= lastNonLeaf {
		left := 2 * index
		right := 2*index + 1
		target := index
		if h.isSiftUp(left, target) {
			target = left
		}
		if h.isSiftUp(right, target) {
			target = right
		}
		if target != index {
			h.swap(index, target)
		} else {
			break
		}
		index = target
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
