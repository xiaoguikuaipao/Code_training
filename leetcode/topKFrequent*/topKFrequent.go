package topKFrequent

type kv struct {
	key   int
	value int
}

type MinHeap struct {
	heap []kv
}

type MaxHeap struct {
	heap []int
}

func (h *MinHeap) Pop() kv {
	pop := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.siftDown(0)
	return pop
}

func (h *MinHeap) Push(x kv) {
	h.heap = append(h.heap, x)
	h.siftUp(len(h.heap) - 1)
}

func (h *MinHeap) siftDown(index int) {
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	swapIndex := index

	//choose the smaller to swap
	if leftChild < len(h.heap) && h.heap[leftChild].value < h.heap[swapIndex].value {
		swapIndex = leftChild
	}

	if rightChild < len(h.heap) && h.heap[rightChild].value < h.heap[swapIndex].value {
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
	if parent >= 0 && h.heap[index].value < h.heap[parent].value {
		h.swap(parent, index)
		h.siftUp(parent)
	}
}

func (h *MinHeap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func topKFrequent(nums []int, k int) []int {
	minHeap := &MinHeap{}
	table := make(map[int]int)
	ret := make([]int, 0)
	for _, e := range nums {
		table[e]++
	}
	for key, value := range table {
		minHeap.Push(kv{key, value})
		if len(minHeap.heap) > k {
			minHeap.Pop()
		}
	}
	for _, e := range minHeap.heap {
		ret = append(ret, e.key)
	}
	return ret
}
