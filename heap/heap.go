// This package provides implementation for heap data structure with generics support.
package heap

import (
	"cmp"
)

type Heap[T any] struct {
	data []T
	less func(a, b T) bool
}

// Initializes an empty max heap.
// Works with default built in types.
func NewMaxHeap[T cmp.Ordered]() *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0, 10),
		less: func(a, b T) bool { return a > b },
	}
}

// Initializes an empty min heap.
// Works with default built in types.
func NewMinHeap[T cmp.Ordered]() *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0, 10),
		less: func(a, b T) bool { return a < b },
	}
}

// Initializes an empty heap.
// Works with any custom type as defined by the user.
// Takes a comparator function that defines the behaviour of the heap.
// To make a min heap, use a < b comparison.
// To make a max heap, use a > b comparison.
func NewHeapWithFunc[T any](comparator func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0, 10),
		less: comparator,
	}
}

func (h *Heap[T]) Push(val T) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
}

func (h *Heap[T]) Pop() {
	if len(h.data) == 0 {
		return
	}
	lastIndex := len(h.data) - 1
	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]
	h.heapifyDown(0)
}

func (h *Heap[T]) Top() T {
	if len(h.data) == 0 {
		return *new(T) // return zero value for the target type
	}

	return h.data[0]
}

func (h *Heap[T]) Size() int {
	return len(h.data)
}

func (h *Heap[T]) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap[T]) heapifyUp(index int) {
	currentIndex := index
	for currentIndex > 0 {
		parentIndex := (currentIndex - 1) / 2
		if !h.less(h.data[currentIndex], h.data[parentIndex]) {
			break
		}

		h.data[currentIndex], h.data[parentIndex] = h.data[parentIndex], h.data[currentIndex]

		currentIndex = parentIndex
	}
}

func (h *Heap[T]) heapifyDown(index int) {
	currentIndex := index

	for currentIndex < len(h.data) {
		smallerIndex := currentIndex
		leftChildIndex := 2*currentIndex + 1
		rightChildIndex := 2*currentIndex + 2

		if leftChildIndex < len(h.data) && h.less(h.data[leftChildIndex], h.data[smallerIndex]) {
			smallerIndex = leftChildIndex
		}

		if rightChildIndex < len(h.data) && h.less(h.data[rightChildIndex], h.data[smallerIndex]) {
			smallerIndex = rightChildIndex
		}

		if smallerIndex == currentIndex {
			break
		}

		h.data[smallerIndex], h.data[currentIndex] = h.data[currentIndex], h.data[smallerIndex]

		currentIndex = smallerIndex
	}
}
