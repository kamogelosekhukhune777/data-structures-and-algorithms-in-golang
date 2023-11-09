package minHeap

import "fmt"

type MinHeap struct {
	heap []int
}

func NewMinHeap() *MinHeap {
	return new(MinHeap)
}

func (h *MinHeap) Insert(key int) {
	h.heap = append(h.heap, key)
	h.heapifyUp(len(h.heap) - 1)
}
func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.heap[index] >= h.heap[parentIndex] {
			break
		}
		h.heap[index], h.heap[parentIndex] = h.heap[parentIndex], h.heap[index]
		index = parentIndex
	}
}

func (h *MinHeap) ExtractMin() (int, error) {
	if len(h.heap) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	min := h.heap[0]
	lastIndex := len(h.heap) - 1
	h.heap[0] = h.heap[lastIndex]
	h.heap = h.heap[:lastIndex]
	h.heapifyDown(0)

	return min, nil
}
func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.heap) - 1
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2
	smallestIndex := leftChildIndex
	if leftChildIndex <= lastIndex && h.heap[leftChildIndex] < h.heap[lastIndex] {
		smallestIndex = leftChildIndex
	}
	if rightChildIndex <= lastIndex && h.heap[rightChildIndex] < h.heap[lastIndex] {
		smallestIndex = leftChildIndex
	}
	if smallestIndex != index {
		h.heap[index], h.heap[smallestIndex] = h.heap[smallestIndex], h.heap[index]
		h.heapifyDown(smallestIndex)
	}
}
