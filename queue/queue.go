package queue

import "errors"

type queue[T any] struct {
	internalSlice     []T
	internalSliceSize int
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{
		internalSlice:     make([]T, 0, 10),
		internalSliceSize: 0,
	}
}

func (q *queue[T]) Push(val T) {
	q.internalSlice = append(q.internalSlice, val)
	q.internalSliceSize++
}

func (q *queue[T]) Pop() error {
	if q.internalSliceSize == 0 {
		return errors.New("queue is empty")
	}
	q.internalSlice = q.internalSlice[:q.internalSliceSize]
	q.internalSliceSize--
	return nil
}

func (q *queue[T]) Front() (T, error) {
	if q.internalSliceSize != 0 {
		return q.internalSlice[0], nil
	}

	return *new(T), errors.New("queue is empty")
}

func (q *queue[T]) Back() (T, error) {
	if q.internalSliceSize != 0 {
		return q.internalSlice[q.internalSliceSize-1], nil
	}

	return *new(T), errors.New("queue is empty")
}

func (q *queue[T]) IsEmpty() bool {
	return q.internalSliceSize == 0
}
