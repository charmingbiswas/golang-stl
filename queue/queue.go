// This package implements a deque.
package queue

import (
	"container/list"
)

type Queue[T any] struct {
	data *list.List
}

// Initializes an empty queue.
// Works with any generic data type.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: list.New(),
	}
}

// Add an element to the front of the queue.
func (this *Queue[T]) PushFront(val T) {
	this.data.PushFront(val)
}

// Add an element to the end of the queue.
func (this *Queue[T]) PushBack(val T) {
	this.data.PushBack(val)
}

// Remove an element from the front of the queue.
func (this *Queue[T]) PopFront() {
	if this.data.Len() == 0 {
		return
	}
	this.data.Remove(this.data.Front())
}

// Remove an element from the back of the queue.
func (this *Queue[T]) PopBack() {
	if this.data.Len() == 0 {
		return
	}
	this.data.Remove(this.data.Back())
}

// Returns the first element in the queue.
func (this *Queue[T]) Front() T {
	if this.data.Len() == 0 {
		return *new(T)
	}
	return this.data.Front().Value.(T)
}

// Returns the last element in the queue.
func (this *Queue[T]) Back() T {
	if this.data.Len() == 0 {
		return *new(T)
	}
	return this.data.Back().Value.(T)
}

// Checks if the queue is empty.
// Return boolean.
func (this *Queue[T]) IsEmpty() bool {
	return this.data.Len() == 0
}
