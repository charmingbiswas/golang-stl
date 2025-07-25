package heap

import (
	"cmp"
)

type heap[T any] struct {
	data []T
	less func(a, b T) bool
}

func NewMaxHeap[T cmp.Ordered]() *heap[T] {
	return &heap[T]{
		data: make([]T, 0, 10),
		less: func(a, b T) bool { return a > b },
	}
}

func NewMinHeap[T cmp.Ordered]() *heap[T] {
	return &heap[T]{
		data: make([]T, 0, 10),
		less: func(a, b T) bool { return a < b },
	}
}

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

func (this *heap[T]) Pop() bool {
	if len(this.data) == 0 {
		return false
	}
	lastIndex := len(this.data) - 1
	this.data[0] = this.data[lastIndex]
	this.data = this.data[:lastIndex]
	this.heapifyDown(0)
	return true
}

func (this *heap[T]) Top() (T, bool) {
	if len(this.data) == 0 {
		return *new(T), false
	}

	return this.data[0], true
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
