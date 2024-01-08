package heap_dataStructure

type MinHeap struct {
	heap []int
}
type MaxHeap struct {
	heap []int
}

func (h *MinHeap) Build(data []int) {
	h.heap = append(h.heap, data...)
	lastNotLeaf := len(h.heap)/2 - 1

	//from lastNotLeaf, from down to up, from right to left...
	for i := lastNotLeaf; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *MinHeap) Pop() int {
	pop := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.siftDown(0)
	return pop
}

func (h *MinHeap) Push(x int) {
	h.heap = append(h.heap, x)
	h.siftUp(len(h.heap) - 1)
}

func (h *MinHeap) siftDown(index int) {
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	swapIndex := index

	//choose the smaller to swap
	if leftChild < len(h.heap) && h.heap[leftChild] < h.heap[swapIndex] {
		swapIndex = leftChild
	}

	if rightChild < len(h.heap) && h.heap[rightChild] < h.heap[swapIndex] {
		swapIndex = rightChild
	}

	if index != swapIndex {
		h.swap(index, swapIndex)
		//recurse
		h.siftDown(swapIndex)
	}

}

func (h *MinHeap) siftUp(index int) {
	parent := (index - 1) / 2
	if parent >= 0 && h.heap[index] < h.heap[parent] {
		h.swap(parent, index)
		h.siftUp(parent)
	}
}

func (h *MinHeap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}
