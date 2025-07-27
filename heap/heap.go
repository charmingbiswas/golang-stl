// This package provides implementation for heap data structure with generics support.
package heap

import (
	"cmp"
)

type heap[T any] struct {
	data []T
	less func(a, b T) bool
}

// Initializes an empty max heap.
// Works with default built in types.
func NewMaxHeap[T cmp.Ordered]() *heap[T] {
	return &heap[T]{
		data: make([]T, 0, 10),
		less: func(a, b T) bool { return a > b },
	}
}

// Initializes an empty min heap.
// Works with default built in types.
func NewMinHeap[T cmp.Ordered]() *heap[T] {
	return &heap[T]{
		data: make([]T, 0, 10),
		less: func(a, b T) bool { return a < b },
	}
}

// Initializes an empty heap.
// Works with any custom type as defined by the user.
// Takes a comparator function that defines the behaviour of the heap.
// To make a min heap, use a < b comparison.
// To make a max heap, use a > b comparison.
func NewHeapWithFunc[T any](comparator func(a, b T) bool) *heap[T] {
	return &heap[T]{
		data: make([]T, 0, 10),
		less: comparator,
	}
}

func (this *heap[T]) Push(val T) {
	this.data = append(this.data, val)
	this.heapifyUp(len(this.data) - 1)
}

func (this *heap[T]) Pop() {
	if len(this.data) == 0 {
		return
	}
	lastIndex := len(this.data) - 1
	this.data[0] = this.data[lastIndex]
	this.data = this.data[:lastIndex]
	this.heapifyDown(0)
}

func (this *heap[T]) Top() T {
	if len(this.data) == 0 {
		return *new(T) // return zero value for the target type
	}

	return this.data[0]
}

func (this *heap[T]) Size() int {
	return len(this.data)
}

func (this *heap[T]) IsEmpty() bool {
	return len(this.data) == 0
}

func (this *heap[T]) heapifyUp(index int) {
	currentIndex := index
	for currentIndex > 0 {
		parentIndex := (currentIndex - 1) / 2
		if !this.less(this.data[currentIndex], this.data[parentIndex]) {
			break
		}

		this.data[currentIndex], this.data[parentIndex] = this.data[parentIndex], this.data[currentIndex]

		currentIndex = parentIndex
	}
}

func (this *heap[T]) heapifyDown(index int) {
	currentIndex := index

	for currentIndex < len(this.data) {
		smallerIndex := currentIndex
		leftChildIndex := 2*currentIndex + 1
		rightChildIndex := 2*currentIndex + 2

		if leftChildIndex < len(this.data) && this.less(this.data[leftChildIndex], this.data[smallerIndex]) {
			smallerIndex = leftChildIndex
		}

		if rightChildIndex < len(this.data) && this.less(this.data[rightChildIndex], this.data[smallerIndex]) {
			smallerIndex = rightChildIndex
		}

		if smallerIndex == currentIndex {
			break
		}

		this.data[smallerIndex], this.data[currentIndex] = this.data[currentIndex], this.data[smallerIndex]

		currentIndex = smallerIndex
	}
}
