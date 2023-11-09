package maxHeap

import "fmt"

type MaxHeap struct {
	heap []int
}

func NewMaxHeap() *MaxHeap {
	return new(MaxHeap)
}

func (h *MaxHeap) Insert(key int) {
	h.heap = append(h.heap, key)
	h.heapifyUp(len(h.heap) - 1)
}
func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.heap[parentIndex] >= h.heap[index] {
			break
		}
		h.heap[index], h.heap[parentIndex] = h.heap[parentIndex], h.heap[index]
		index = parentIndex
	}
}

//ExtractMax removes the root node(node(key)with highest priority)
func (h *MaxHeap) ExtractMax() (int, error) {
	if len(h.heap) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	max := h.heap[0]
	lastIndex := len(h.heap) - 1
	h.heap[0] = h.heap[lastIndex]
	h.heap = h.heap[:lastIndex]
	h.heapifyDown(0)

	return max, nil
}
func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.heap) - 1
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2
	largestIndex := index

	if leftChildIndex <= lastIndex && h.heap[leftChildIndex] > h.heap[largestIndex] {
		largestIndex = leftChildIndex
	}
	if rightChildIndex <= lastIndex && h.heap[rightChildIndex] > h.heap[largestIndex] {
		largestIndex = rightChildIndex
	}
	if largestIndex != index {
		h.heap[index], h.heap[largestIndex] = h.heap[largestIndex], h.heap[index]
		h.heapifyDown(largestIndex)
	}
}
