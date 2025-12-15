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
func (q *Queue[T]) PushFront(val T) {
	q.data.PushFront(val)
}

// Add an element to the end of the queue.
func (q *Queue[T]) PushBack(val T) {
	q.data.PushBack(val)
}

// Remove an element from the front of the queue.
func (q *Queue[T]) PopFront() {
	if q.data.Len() == 0 {
		return
	}
	q.data.Remove(q.data.Front())
}

// Remove an element from the back of the queue.
func (q *Queue[T]) PopBack() {
	if q.data.Len() == 0 {
		return
	}
	q.data.Remove(q.data.Back())
}

// Returns the first element in the queue.
func (q *Queue[T]) Front() T {
	if q.data.Len() == 0 {
		return *new(T)
	}
	return q.data.Front().Value.(T)
}

// Returns the last element in the queue.
func (q *Queue[T]) Back() T {
	if q.data.Len() == 0 {
		return *new(T)
	}
	return q.data.Back().Value.(T)
}

// Checks if the queue is empty.
// Return boolean.
func (q *Queue[T]) IsEmpty() bool {
	return q.data.Len() == 0
}

func (q *Queue[T]) Size() int {
	return q.data.Len()
}
