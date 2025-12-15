// This package implements a deque.
package queue

import (
	"container/list"
	"errors"
)

var (
	ErrEmptyQueue = errors.New("queue is empty")
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

// Adds an element to the front of the queue.
func (q *Queue[T]) PushFront(val T) {
	q.data.PushFront(val)
}

// Adds an element to the end of the queue.
func (q *Queue[T]) PushBack(val T) {
	q.data.PushBack(val)
}

// Removes an element from the front of the queue.
// Returns an error if queue is empty.
func (q *Queue[T]) PopFront() error {
	if q.data.Len() == 0 {
		return ErrEmptyQueue
	}
	q.data.Remove(q.data.Front())
	return nil
}

// Removes an element from the back of the queue.
// Returns an error is queue is empty.
func (q *Queue[T]) PopBack() error {
	if q.data.Len() == 0 {
		return ErrEmptyQueue
	}
	q.data.Remove(q.data.Back())
	return nil
}

// Returns the first element in the queue.
// Retuns an error if queue is empty.
func (q *Queue[T]) Front() (T, error) {
	if q.data.Len() == 0 {
		return *new(T), ErrEmptyQueue
	}
	return q.data.Front().Value.(T), nil
}

// Returns the last element in the queue.
// Returns an error if queue is empty.
func (q *Queue[T]) Back() (T, error) {
	if q.data.Len() == 0 {
		return *new(T), ErrEmptyQueue
	}
	return q.data.Back().Value.(T), nil
}

// Checks if the queue is empty.
// Returns boolean.
func (q *Queue[T]) IsEmpty() bool {
	return q.data.Len() == 0
}

// Returns the current size of the queue.
func (q *Queue[T]) Size() int {
	return q.data.Len()
}
